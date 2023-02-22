package service_test

import (
	"context"
	"net"
	"testing"

	pb "github.com/rabbice/fixturesbook/pb"
	"github.com/rabbice/fixturesbook/sample"
	"github.com/rabbice/fixturesbook/serializer"
	"github.com/rabbice/fixturesbook/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientCreateFixture(t *testing.T) {
	t.Parallel()

	fixtureStore := service.NewInMemoryFixtureStore()
	serverAddress := startTestFixtureServer(t, fixtureStore)
	fixtureClient := newTestFixtureClient(t, serverAddress)

	fixture := sample.NewFixture()
	expectedID := fixture.ID
	req := &pb.CreateFixtureRequest{
		Fixture: fixture,
	}

	res, err := fixtureClient.CreateFixture(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// check that the fixture is saved to the store
	other, err := fixtureStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that the saved fixture is the same as the one we send
	requireSameFixture(t, fixture, other)
}

func startTestFixtureServer(t *testing.T, fixtureStore service.FixtureStore) string {
	fixtureServer := service.NewFixtureServer(fixtureStore)

	grpcServer := grpc.NewServer()
	pb.RegisterFixtureServiceServer(grpcServer, fixtureServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return listener.Addr().String()
}

func newTestFixtureClient(t *testing.T, serverAddress string) pb.FixtureServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return pb.NewFixtureServiceClient(conn)
}

func requireSameFixture(t *testing.T, fixture1 *pb.Fixture, fixture2 *pb.Fixture) {
	json1, err := serializer.ProtobufToJSON(fixture1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(fixture2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}

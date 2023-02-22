package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/rabbice/fixturesbook/pb"
	"github.com/rabbice/fixturesbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func createFixture(fixtureClient pb.FixtureServiceClient) {
	fixture := sample.NewFixture()
	fixture.ID = ""
	req := &pb.CreateFixtureRequest{
		Fixture: fixture,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := fixtureClient.CreateFixture(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("fixture already exists")
		} else {
			log.Fatal("cannot create fixture: ", err)
		}
		return
	}
	log.Printf("created fixture with id: %s", res.Id)
}

func searchFixture(fixtureClient pb.FixtureServiceClient, filter *pb.Filter) {
	log.Printf("search filter: %v", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchFixtureRequest{Filter: filter}
	stream, err := fixtureClient.SearchFixture(ctx, req)
	if err != nil {
		log.Fatal("cannot search fixture: ", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		fixture := res.Fixture
		log.Print("- found: ", fixture.GetID())
		log.Print("- found: ", fixture.GetHometeam())
		log.Print("- found: ", fixture.GetAwayteam())
		log.Print("- found: ", fixture.GetPitch())
		log.Print("- found: ", fixture.GetOfficial())
	}
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dialing server: %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	fixtureClient := pb.NewFixtureServiceClient(conn)
	for i := 0; i < 10; i++ {
		createFixture(fixtureClient)
	}

	filter := &pb.Filter{
		Date: "2021-01-22",
	}

	searchFixture(fixtureClient, filter)
}

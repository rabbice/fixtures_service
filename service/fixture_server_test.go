package service_test

import (
	"context"
	"testing"

	pb "github.com/rabbice/fixturesbook/pb"
	"github.com/rabbice/fixturesbook/sample"
	"github.com/rabbice/fixturesbook/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	fixtureNoID := sample.NewFixture()
	fixtureNoID.ID = " "

	fixtureInvalidID := sample.NewFixture()
	fixtureInvalidID.ID = "invalid-uuid"

	fixtureDuplicateID := sample.NewFixture()
	storeDuplicateID := service.NewInMemoryFixtureStore()
	err := storeDuplicateID.Save(fixtureDuplicateID)
	require.Nil(t, err)

	testCases := []struct {
		name    string
		fixture *pb.Fixture
		store   service.FixtureStore
		code    codes.Code
	}{
		{
			name:    "success_with_id",
			fixture: sample.NewFixture(),
			store:   service.NewInMemoryFixtureStore(),
			code:    codes.OK,
		},
		{
			name:    "success_no_id",
			fixture: fixtureNoID,
			store:   service.NewInMemoryFixtureStore(),
			code:    codes.InvalidArgument,
		},
		{
			name:    "failure_invalid_id",
			fixture: fixtureInvalidID,
			store:   service.NewInMemoryFixtureStore(),
			code:    codes.InvalidArgument,
		},
		{
			name:    "failure_duplicate_id",
			fixture: fixtureDuplicateID,
			store:   storeDuplicateID,
			code:    codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateFixtureRequest{
				Fixture: tc.fixture,
			}

			server := service.NewFixtureServer(tc.store)
			res, err := server.CreateFixture(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.fixture.ID) > 0 {
					require.Equal(t, tc.fixture.ID, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}

}

package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	pb "github.com/rabbice/fixturesbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FixtureServer is the server that provides fixture services
type FixtureServer struct {
	pb.UnimplementedFixtureServiceServer
	fixtureStore FixtureStore
}

// NewFixtureServer returns a new FixtureServer
func NewFixtureServer(fixtureStore FixtureStore) *FixtureServer {
	return &FixtureServer{
		fixtureStore: fixtureStore,
	}
}

func (s *FixtureServer) CreateFixture(
	ctx context.Context,
	req *pb.CreateFixtureRequest,
) (*pb.CreateFixtureResponse, error) {
	fixture := req.GetFixture()
	log.Printf("received a create fixture request with id: %s", fixture.ID)

	if len(fixture.ID) > 0 {
		_, err := uuid.Parse(fixture.ID)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "fixture Id is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new fixture ID: %v", err)
		}
		fixture.ID = id.String()
	}

	// heavy processing
	//time.Sleep(6 * time.Second)

	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline to client exceeded")
	}

	// save fixture to in-memory store
	err := s.fixtureStore.Save(fixture)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save fixture to the store: %v", err)
	}
	log.Printf("saved fixture with id: %s", fixture.ID)

	res := &pb.CreateFixtureResponse{
		Id: fixture.ID,
	}
	return res, nil
}

func (server *FixtureServer) SearchFixture(req *pb.SearchFixtureRequest, stream pb.FixtureService_SearchFixtureServer) error {
	filter := req.GetFilter()
	log.Printf("received search fixture request with filter: %v", filter)

	err := server.fixtureStore.Search(
		stream.Context(),
		filter,
		func(fixture *pb.Fixture) error {
			res := &pb.SearchFixtureResponse{Fixture: fixture}

			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("sent fixture with id: %s", fixture.GetID())
			return nil
		},
	)

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}

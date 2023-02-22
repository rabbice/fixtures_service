package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/copier"

	pb "github.com/rabbice/fixturesbook/pb"
)

var ErrAlreadyExists = errors.New("record already exists")

type FixtureStore interface {
	// saves fixture to the store
	Save(fixture *pb.Fixture) error
	// find fixture by ID
	Find(id string) (*pb.Fixture, error)
	// search for fixture that matches certain parameters
	Search(ctx context.Context, filter *pb.Filter, found func(fixture *pb.Fixture) error) error
}

type InMemoryFixtureStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Fixture
}

func NewInMemoryFixtureStore() *InMemoryFixtureStore {
	return &InMemoryFixtureStore{
		data: make(map[string]*pb.Fixture),
	}
}

func (store *InMemoryFixtureStore) Save(fixture *pb.Fixture) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[fixture.ID] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &pb.Fixture{}
	err := copier.Copy(other, fixture)
	if err != nil {
		return fmt.Errorf("cannot copy fixture data: %v", err)
	}

	store.data[other.ID] = other
	return nil
}

// Find finds a fixture by ID
func (store *InMemoryFixtureStore) Find(id string) (*pb.Fixture, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	fixture := store.data[id]
	if fixture == nil {
		return nil, nil
	}

	return deepCopy(fixture)
}

func (store *InMemoryFixtureStore) Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(fixture *pb.Fixture) error,
) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, fixture := range store.data {
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is canceled")
			return nil
		}

		if isQualified(filter, fixture) {
			other, err := deepCopy(fixture)
			if err != nil {
				return err
			}

			err = found(other)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func deepCopy(fixture *pb.Fixture) (*pb.Fixture, error) {
	other := &pb.Fixture{}

	err := copier.Copy(other, fixture)
	if err != nil {
		return nil, fmt.Errorf("cannot copy fixture data: %w", err)
	}

	return other, nil
}

func isQualified(filter *pb.Filter, fixture *pb.Fixture) bool {
	if fixture.Time.Date != filter.GetDate() {
		return false
	}
	return true
}

package comment

import (
	"context"
	"fmt"
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment Body
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Struct where all our logic will be built on top of
type Service struct {
	Store Store
}

// Returns pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a commnet")
	comment, err := s.Store.GetComment(ctx, "Here")
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return comment, nil
}

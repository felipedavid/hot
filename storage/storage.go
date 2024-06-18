package storage

import (
	"context"

	"github.com/felipedavid/hot/types"
)

type Storage interface {
	GetUser(ctx context.Context, id string) (*types.User, error)
}

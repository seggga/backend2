package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/seggga/backend2/internal/entity"
)

// Repo is interface to interact with particular repository
type Repo interface {
	CreateUser(ctx context.Context, u entity.User) (*entity.User, error)
	CreateGroup(ctx context.Context, u entity.Group) (*entity.Group, error)
	ReadUser(ctx context.Context, uid uuid.UUID) (*entity.User, error)
	ReadGroup(ctx context.Context, uid uuid.UUID) (*entity.Group, error)
	AddToGroup(ctx context.Context, uid, gid uuid.UUID) error
	RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error
	SearchUser(ctx context.Context, name string, ids ...uuid.UUID) ([]uuid.UUID, error)
	SearchGroup(ctx context.Context, name string, ids ...uuid.UUID) ([]uuid.UUID, error)
}

// Storage implements main logic using Repo methods
type Storage struct {
	Repo Repo
}

// NewStorage creates Storage instance
func NewStorage(repo Repo) *Storage {
	return &Storage{
		Repo: repo,
	}
}

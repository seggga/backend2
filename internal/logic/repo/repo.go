package repo

import (
	"context"
	"os/user"

	"github.com/google/uuid"
	"github.com/seggga/backend2/internal/entity"
)

// Repo is interface to interact with particular repository
type Repo interface {
	CreateUser(ctx context.Context, u entity.User) (*entity.User, error)
	CreateGroup(ctx context.Context, u entity.Group) (*entity.Group, error)
	Read(ctx context.Context, uid uuid.UUID) (*user.User, error)
	AddToGroup(ctx context.Context, uid uuid.UUID) error
	RemoveFromGroup(ctx context.Context, uid uuid.UUID) error
	SearchUser(ctx context.Context, s string) (chan user.User, error)
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

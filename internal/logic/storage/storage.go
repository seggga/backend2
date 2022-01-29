package storage

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seggga/backend2/internal/entity"
)

// Repo is interface to interact with particular storage repository
type Repo interface {
	CreateUser(ctx context.Context, u entity.User) (*entity.User, error)
	CreateGroup(ctx context.Context, u entity.Group) (*entity.Group, error)
	ReadUser(ctx context.Context, uid uuid.UUID) (*entity.User, error)
	ReadGroup(ctx context.Context, uid uuid.UUID) (*entity.Group, error)
	AddToGroup(ctx context.Context, uid, gid uuid.UUID) error
	RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error
	SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error)
	SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error)
}

// DB implements main logic using Repo methods
type DB struct {
	repo Repo
}

// NewDB creates Storage instance
func NewDB(repo Repo) *DB {
	return &DB{
		repo: repo,
	}
}

// CreateUser adds ID to the passed user data and calls repo's CreateUser method
func (db *DB) CreateUser(ctx context.Context, u entity.User) (*entity.User, error) {
	u.ID = uuid.New()
	newUser, err := db.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("error creating new user: %w", err)
	}
	return newUser, nil
}

// CreateGroup adds ID to the passed group data and calls repo's CreateGroup method
func (db *DB) CreateGroup(ctx context.Context, g entity.Group) (*entity.Group, error) {
	g.ID = uuid.New()
	newGroup, err := db.repo.CreateGroup(ctx, g)
	if err != nil {
		return nil, fmt.Errorf("error creating new user: %w", err)
	}
	return newGroup, nil
}

// AddToGroup calls same function from repo
func (db *DB) AddToGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return db.repo.AddToGroup(ctx, uid, gid)
}

// RemoveFromGroup calls same function from repo
func (db *DB) RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error {
	return db.repo.RemoveFromGroup(ctx, uid, gid)
}

// SearchUser ...
func (db *DB) SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error) {
	users, err := db.repo.SearchUser(ctx, name, gids)
	if err != nil {
		return nil, fmt.Errorf("error searching users: %w", err)
	}
	return users, nil
}

// SearchGroup ...
func (db *DB) SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error) {
	groups, err := db.repo.SearchGroup(ctx, name, uids)
	if err != nil {
		return nil, fmt.Errorf("error searching groups: %w", err)
	}
	return groups, nil
}

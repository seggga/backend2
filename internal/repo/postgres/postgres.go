package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/seggga/backend2/internal/entity"

	"log"
)

func main() {

	db, err := sql.Open("postgres", "user=adm password=pass dbname=ung sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	age := 21
	rows, err := db.Query("SELECT city FROM streams WHERE id=$1", age)

	fmt.Println(rows)

}

type PGRepo struct{}

func InitDB() {}

func (pg *PGRepo) CreateUser(ctx context.Context, u entity.User) (*entity.User, error)
func (pg *PGRepo) CreateGroup(ctx context.Context, u entity.Group) (*entity.Group, error)
func (pg *PGRepo) ReadUser(ctx context.Context, uid uuid.UUID) (*entity.User, error)
func (pg *PGRepo) ReadGroup(ctx context.Context, uid uuid.UUID) (*entity.Group, error)
func (pg *PGRepo) AddToGroup(ctx context.Context, uid, gid uuid.UUID) error
func (pg *PGRepo) RemoveFromGroup(ctx context.Context, uid, gid uuid.UUID) error
func (pg *PGRepo) SearchUser(ctx context.Context, name string, gids []uuid.UUID) ([]entity.User, error)
func (pg *PGRepo) SearchGroup(ctx context.Context, name string, uids []uuid.UUID) ([]entity.Group, error)

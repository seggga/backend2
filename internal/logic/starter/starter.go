package starter

import (
	"context"

	"sync"

	"github.com/seggga/backend2/internal/logic/storage"
)

type App struct {
	db *storage.DB
}

func NewApp(repo storage.Repo) *App {
	a := &App{
		db: storage.NewDB(repo),
	}
	return a
}

type APIServer interface {
	Start(db *storage.DB)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs APIServer) {
	defer wg.Done()
	hs.Start(a.db)
	<-ctx.Done()
	hs.Stop()
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/seggga/backend2/internal/api/handler"
	"github.com/seggga/backend2/internal/api/server"
	"github.com/seggga/backend2/internal/logic/starter"
	"github.com/seggga/backend2/internal/logic/storage"
	"github.com/seggga/backend2/internal/repo/postgres"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// "postgresql://user:password@host:port/dbname"
	connString := `postgresql://appuser:appp@$$w0rd@127.0.0.1:5432/app_db`
	DB, err := postgres.NewDB(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}

	a := starter.NewApp(DB)
	h := handler.NewRouter(storage.NewDB(DB))
	srv := server.NewServer(":8000", h)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}

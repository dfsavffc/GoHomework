package main

import (
	"context"
	"errors"
	"github.com/dfsavffc/GoHomework/task2/server/internal/app"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 5 * time.Second
const host = 8000

func main() {

	server := app.NewServer(host)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		log.Printf("start listen on: %d\n", host)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error in listen: %s\n", err)
			return err
		}
		log.Println("server stopped")
		return nil
	})
	group.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("error in shutdown: %s\n", err)
			return err
		}
		log.Println("server shutdown")
		return nil
	})
	err := group.Wait()
	if err != nil {
		log.Fatalf("error in wait: %s\n", err)
		return
	}
}

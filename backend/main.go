package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/giornetta/devcv/repository/inmem"

	"github.com/giornetta/devcv/devcv"

	"github.com/giornetta/devcv/cfg"
	"github.com/giornetta/devcv/server"

	"github.com/giornetta/devcv/auth"

	"github.com/giornetta/devcv/developers"

	"github.com/giornetta/devcv/repository/postgres"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	c, err := cfg.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	authSvc := auth.New(c.JWTSecret)

	var repo devcv.DeveloperRepository

	if c.DBType == "postgres" {
		db, err := postgres.Connect(c.DBHost, c.DBPort, c.DBName, c.DBUser, c.DBPassword)
		if err != nil {
			log.Fatalf("could not open db: %v", err)
		}

		repo = postgres.NewDeveloperRepository(db)
	} else {
		repo = inmem.NewDeveloperRepository()
	}

	developersSvc := developers.New(repo, authSvc)

	srv := server.New(c.HTTPPort, developersSvc, authSvc)

	go func() {
		log.Println("Starting HTTP server...")
		srv.ListenAndServe()
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	<-sig
	log.Println("Shutting down server...")
	srv.Shutdown(ctx)
	db.Close()
}

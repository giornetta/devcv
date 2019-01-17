package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/giornetta/devcv/auth"

	"github.com/giornetta/devcv/proto"

	"github.com/giornetta/devcv/developers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/giornetta/devcv/repository"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	var (
		dbHost     = "db"
		dbPort     = 5432
		dbUser     = "postgres-dev"
		dbPassword = "password"
		dbName     = "dev"

		grpcPort = ":3001"
		httpPort = ":3000"
	)

	authService := auth.New("topsecret")

	db, err := repository.NewDB(dbHost, dbPort, dbName, dbUser, dbPassword)
	if err != nil {
		log.Fatalf("could not open db: %v", err)
	}

	repo := repository.NewDevelopers(db)
	var developersService proto.DeveloperServiceServer
	{
		developersService = developers.New(repo)
		developersService = developers.NewAuthenticator(developersService, authService)
	}

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterDeveloperServiceServer(grpcServer, developersService)

	go func() {
		log.Println("Starting GRPC server...")
		grpcServer.Serve(lis)
	}()

	// HTTP
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := proto.RegisterDeveloperServiceHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts); err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    httpPort,
		Handler: allowCORS(mux),
	}

	go func() {
		log.Println("Starting HTTP server...")
		httpServer.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c
	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	httpServer.Shutdown(ctx)
	db.Close()

	os.Exit(0)
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				headers := []string{"Content-Type", "Accept", "Authorization"}
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
				methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

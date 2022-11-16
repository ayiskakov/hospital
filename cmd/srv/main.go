package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/khanfromasia/hospital/internal/config"
	handler "github.com/khanfromasia/hospital/internal/delivery/http"
	"github.com/khanfromasia/hospital/internal/pkg/pgxpool"
	"github.com/khanfromasia/hospital/internal/storage"
)

func main() {

	ctx := context.Background()

	if err := config.ReadConfigYML("./config.yml"); err != nil {
		log.Fatalf("failed to read config fail %s", err.Error())
	}

	cfg := config.Get()

	m, err := migrate.New("file://migrations", "postgres://postgres:secret@localhost:15432/hospital?sslmode=disable")
	if err != nil {
		log.Fatalln("migrations: while connecting", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("migrations: while migration", err)
	}

	pool, err := pgxpool.NewPool(ctx, cfg.Database.Dsn)

	if err != nil {
		log.Println("Failed to open a postgres pool: ", err.Error())
		return
	}

	defer pool.Close()

	stg := storage.NewStorage(pool)
	hnd := handler.NewHandler(stg)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Rest.Host, cfg.Rest.Port))

	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	srv := &http.Server{
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 40 << 20, // 1 MB
		Handler:        corsMiddleware(hnd.SetupRoutes()),
	}

	log.Println("HTTP server is running on ", fmt.Sprintf("%s:%d", cfg.Rest.Host, cfg.Rest.Port))
	if err = srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

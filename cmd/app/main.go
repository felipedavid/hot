package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"

	"github.com/felipedavid/hot/handlers"
	"github.com/felipedavid/hot/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	addr := flag.String("addr", ":3000", "http listen addr")
	dbUri := flag.String("db-uri", "mongodb://localhost:27017", "mongodb database uri")
	dbName := flag.String("db-name", "hotel", "database name")
	flag.Parse()

	slog.Info("Connecting to database", "uri", *dbUri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*dbUri))
	if err != nil {
		panic(err)
	}
	database := client.Database(*dbName)

	storage.Init(database)

	server := http.Server{
		Addr:    *addr,
		Handler: handlers.Routes(),
	}

	slog.Info("Running server", "addr", *addr)
	err = server.ListenAndServe()
	panic(err)
}

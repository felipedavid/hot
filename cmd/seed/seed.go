package main

import (
	"context"
	"flag"
	"log/slog"

	"github.com/felipedavid/hot/storage"
	"github.com/felipedavid/hot/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
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
	hotels := []types.Hotel{
		{Name: "La cucura", Location: "Rio de Janeiro"},
		{Name: "Bellucia", Location: "France"},
		{Name: "Casa de la playa", Location: "Mexico"},
		{Name: "La cucura", Location: "Rio de Janeiro"},
	}

	for i := range hotels {
		_ = storage.InsertHotel(context.Background(), &hotels[i])
	}

	rooms := []types.Room{
		{Size: "small", BasePrice: 100, Price: 200, HotelID: hotels[0].ID},
		{Size: "small", BasePrice: 400, Price: 500, HotelID: hotels[0].ID},
		{Size: "small", BasePrice: 200, Price: 300, HotelID: hotels[0].ID},
		{Size: "normal", BasePrice: 300, Price: 400, HotelID: hotels[0].ID},
		{Size: "small", BasePrice: 100, Price: 200, HotelID: hotels[1].ID},
		{Size: "kingsize", BasePrice: 200, Price: 300, HotelID: hotels[1].ID},
		{Size: "small", BasePrice: 400, Price: 500, HotelID: hotels[1].ID},
	}

	for _, room := range rooms {
		_ = storage.InsertRoom(context.Background(), &room)
	}

}

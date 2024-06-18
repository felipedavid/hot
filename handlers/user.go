package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/felipedavid/hot/storage"
	"github.com/felipedavid/hot/types"
)

var stor storage.Storage

func listUsers(w http.ResponseWriter, r *http.Request) error {
	u := types.User{
		FirstName: "James",
		LastName:  "hello",
	}

	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	w.Write(data)

	return nil
}

func getUser(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	u, err := stor.GetUser(context.Background(), id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	w.Write(data)

	return nil
}

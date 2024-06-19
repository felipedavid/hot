package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/felipedavid/hot/storage"
	"github.com/felipedavid/hot/types"
	"golang.org/x/crypto/bcrypt"
)

func listUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := storage.GetUsers(context.Background())
	if err != nil {
		return err
	}

	return writeJSON(w, users)
}

func getUser(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	u, err := storage.GetUser(context.Background(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, u)
}

func createUser(w http.ResponseWriter, r *http.Request) error {
	var params types.CreateUserParams
	err := readJSON(r, &params)
	if err != nil {
		return err
	}

	if !params.Validate() {
		return writeJSON(w, params.CreateUserParamsErrors)
	}

	user, err := types.NewUser(&params)
	if err != nil {
		return err
	}

	err = storage.InsertUser(context.Background(), user)
	if err != nil {
		return err
	}

	return writeJSON(w, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("id")
	err := storage.DeleteUser(context.Background(), userID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return err
		}
		return err
	}

	return writeJSON(w, map[string]string{"msg": fmt.Sprintf("user %s deleted", userID)})
}

func updateUser(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("id")

	var params types.UpdateUserParams
	err := readJSON(r, &params)
	if err != nil {
		return err
	}

	if !params.Validate() {
		return writeJSON(w, params.UpdateUserParamsErrors)
	}

	user, err := storage.GetUser(context.Background(), userID)
	if err != nil {
		return err
	}

	if params.FirstName != nil {
		user.FirstName = *params.FirstName
	}

	if params.LastName != nil {
		user.LastName = *params.LastName
	}

	if params.Password != nil {
		hPassword, err := bcrypt.GenerateFromPassword([]byte(*params.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.HashedPassword = string(hPassword)
	}

	err = storage.UpdateUser(context.Background(), user)
	if err != nil {
		return err
	}

	return writeJSON(w, user)
}

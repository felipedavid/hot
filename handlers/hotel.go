package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/felipedavid/hot/storage"
	"github.com/felipedavid/hot/types"
)

func listHotels(w http.ResponseWriter, r *http.Request) error {
	hotels, err := storage.GetHotels(context.Background())
	if err != nil {
		return err
	}

	return writeJSON(w, hotels)
}

func getHotel(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")

	u, err := storage.GetHotel(context.Background(), id)
	if err != nil {
		return err
	}

	return writeJSON(w, u)
}

func createHotel(w http.ResponseWriter, r *http.Request) error {
	var params types.CreateHotelParams
	err := readJSON(r, &params)
	if err != nil {
		return err
	}

	if !params.Validate() {
		return writeJSON(w, params.HotelParamsErrors)
	}

	hotel := types.NewHotel(&params)

	err = storage.InsertHotel(context.Background(), hotel)
	if err != nil {
		return err
	}

	return writeJSON(w, hotel)
}

func deleteHotel(w http.ResponseWriter, r *http.Request) error {
	hotelID := r.PathValue("id")
	err := storage.DeleteHotel(context.Background(), hotelID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return err
		}
		return err
	}

	return writeJSON(w, map[string]string{"msg": fmt.Sprintf("hotel %s deleted", hotelID)})
}

func updateHotel(w http.ResponseWriter, r *http.Request) error {
	hotelID := r.PathValue("id")

	var params types.UpdateHotelParams
	err := readJSON(r, &params)
	if err != nil {
		return err
	}

	if !params.Validate() {
		return writeJSON(w, params.HotelParamsErrors)
	}

	hotel, err := storage.GetHotel(context.Background(), hotelID)
	if err != nil {
		return err
	}

	if params.Name != nil {
		hotel.Name = *params.Name
	}

	if params.Location != nil {
		hotel.Location = *params.Location
	}

	err = storage.UpdateHotel(context.Background(), hotel)
	if err != nil {
		return err
	}

	return writeJSON(w, hotel)
}

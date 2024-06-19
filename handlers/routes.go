package handlers

import (
	"net/http"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func Routes() http.Handler {
	mux := http.NewServeMux()

	defineUnprotectedRoute(mux, "POST /api/v1/auth", authenticateUser)

	defineRoute(mux, "GET /api/v1/user", listUsers)
	defineRoute(mux, "GET /api/v1/user/{id}", getUser)
	defineRoute(mux, "POST /api/v1/user", createUser)
	defineRoute(mux, "DELETE /api/v1/user/{id}", deleteUser)
	defineRoute(mux, "PATCH /api/v1/user/{id}", updateUser)

	defineRoute(mux, "GET /api/v1/hotel", listHotels)
	defineRoute(mux, "GET /api/v1/hotel/{id}", getHotel)
	defineRoute(mux, "GET /api/v1/hotel/{id}/rooms", getRoomsFromHotel)
	defineRoute(mux, "POST /api/v1/hotel", createHotel)
	defineRoute(mux, "DELETE /api/v1/hotel/{id}", deleteHotel)
	defineRoute(mux, "PATCH /api/v1/hotel/{id}", updateHotel)

	return mux
}

func defineUnprotectedRoute(mux *http.ServeMux, path string, h CustomHandler) {
	mux.HandleFunc(path, centralErrorHandler(h))
}

func defineRoute(mux *http.ServeMux, path string, h CustomHandler) {
	h = Authentication(h)
	mux.HandleFunc(path, centralErrorHandler(h))
}

func centralErrorHandler(h CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			_ = writeJSON(w, map[string]string{
				"error": err.Error(),
			})
		}
	}
}

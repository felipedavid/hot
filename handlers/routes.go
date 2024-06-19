package handlers

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	mux.HandleFunc("GET /api/v1/user", defaultHandler(listUsers))
	mux.HandleFunc("GET /api/v1/user/{id}", defaultHandler(getUser))
	mux.HandleFunc("POST /api/v1/user", defaultHandler(createUser))
	mux.HandleFunc("DELETE /api/v1/user/{id}", defaultHandler(deleteUser))
	mux.HandleFunc("PATCH /api/v1/user/{id}", defaultHandler(updateUser))

	mux.HandleFunc("GET /api/v1/hotel", defaultHandler(listHotels))
	mux.HandleFunc("GET /api/v1/hotel/{id}", defaultHandler(getHotel))
	mux.HandleFunc("POST /api/v1/hotel", defaultHandler(createHotel))
	mux.HandleFunc("DELETE /api/v1/hotel/{id}", defaultHandler(deleteHotel))
	mux.HandleFunc("PATCH /api/v1/hotel/{id}", defaultHandler(updateHotel))

	return mux
}

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func defaultHandler(h CustomHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			_ = writeJSON(w, map[string]string{
				"error": err.Error(),
			})
		}
	}
}

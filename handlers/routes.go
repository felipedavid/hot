package handlers

import (
	"encoding/json"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users", defaultHandler(listUsers))
	mux.HandleFunc("/api/v1/user/{id}", defaultHandler(getUser))

	return mux
}

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func defaultHandler(h CustomHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			e := map[string]string{
				"error": err.Error(),
			}

			data, _ := json.Marshal(e)

			_, _ = w.Write(data)
		}
	}
}

package handlers

import "net/http"

type authenticateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func authenticateUser(w http.ResponseWriter, r *http.Request) error {
	var params authenticateUserParams
	err := readJSON(r, &params)
	if err != nil {
		return err
	}

	return nil
}

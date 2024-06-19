package handlers

import (
	"encoding/json"
	"io"
)

func writeJSON(w io.Writer, msg any) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func readJSON(r io.Reader, dst any) error {
	dec := json.NewDecoder(r)
	return dec.Decode(dst)
}

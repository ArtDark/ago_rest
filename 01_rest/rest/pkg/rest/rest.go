package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteAsJSON(writer http.ResponseWriter, data interface{}) error {
	item, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(item)
	if err != nil {
		log.Print(err)
	}
	return nil
}

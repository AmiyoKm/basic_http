package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Envelop struct {
	Message string `json:"message"`
	Value   any    `json:"value"`
}

func CORSHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "amiyo, Content-Type")
}

func WriteJSON(w http.ResponseWriter, value any) {

	encoder := json.NewEncoder(w)
	err := encoder.Encode(value)
	if err != nil {
		log.Println("error encoding to json", err)
	}
}

func ReadJSON(r *http.Request, value any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(value)
	if err != nil {
		log.Println("error decoding to object", err)
	}
}

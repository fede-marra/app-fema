package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type client struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type credentials struct {
	user     string
	password string
}

type resource struct {
	class  string
	ip     string
	name   string
	domain string
	ports  []uint16
	credentials
}

type jobs struct {
	date       time.Time
	started    bool
	finished   bool
	clientName string
}

func handleClient(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		id := req.URL.Query().Get("id")
		if id == "" {
			clients := []client{}

			jsonResp, err := json.Marshal(clients)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		} else {
			client := client{}
			jsonResp, err := json.Marshal(client)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		}

		return

	case "POST":
		var client client
		err := json.NewDecoder(req.Body).Decode(&client)

		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte{})
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(client)
		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s\n", err)
		}
		w.Write(jsonResp)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte{})
		return
	}
}

func main() {

	http.HandleFunc("/client", handleClient)

	http.ListenAndServe(":8080", nil)
}

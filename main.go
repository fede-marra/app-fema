package main

import (
	"encoding/json"
	"fmt"
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

var clients = make(map[string]client)

func handleClient(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		id := req.URL.Query().Get("Id")
		if id == "" {

			jsonResp, err := json.Marshal(clients)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
			return
		} else {

			client, ok := clients[id]

			if !ok {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			jsonResp, err := json.Marshal(client)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
			return
		}

	case "PUT":
		id := req.URL.Query().Get("Id")

		_, ok := clients[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var client client
		err := json.NewDecoder(req.Body).Decode(&client)

		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte{})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(client)
		if err != nil {
			log.Printf("Error happened in JSON marshal. Err: %s\n", err)
		}

		clients[id] = client
		w.Write(jsonResp)

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

		clients[fmt.Sprint(len(clients)+1)] = client
		w.Write(jsonResp)

	case "DELETE":

		id := req.URL.Query().Get("Id")

		_, ok := clients[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		delete(clients, id)
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte{})
		return
	}
}

func main() {

	// farkim := client{Name: "Farkim", Address: "Pasaje", Phone: "123456", Email: "farkim@farkim.com"}
	// rock := client{Name: "Rock&Fellers", Address: "Orono y Jujuy", Phone: "654321", Email: "rock@rock.com"}
	// labac := client{Name: "Labac", Address: "Nordlink", Phone: "98765", Email: "labac@labac.com"}

	// clientsList["1"] = append(clientsList["1"], farkim)
	// clientsList["2"] = append(clientsList["2"], rock)
	// clientsList["3"] = append(clientsList["3"], labac)

	clients = map[string]client{
		"1": {Name: "Farkim", Address: "Pasaje", Phone: "123456", Email: "farkim@farkim.com"},
		"2": {Name: "Rock&Fellers", Address: "Orono y Jujuy", Phone: "654321", Email: "rock@rock.com"},
		"3": {Name: "Labac", Address: "Nordlink", Phone: "98765", Email: "labac@labac.com"},
	}
	http.HandleFunc("/client", handleClient)

	http.ListenAndServe(":8080", nil)
}

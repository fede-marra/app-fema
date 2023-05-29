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

var clientsList = make(map[string][]client)

func handleClient(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		id := req.URL.Query().Get("Id")
		if id == "" {

			jsonResp, err := json.Marshal(clientsList)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		} else {

			clientFind := clientsList[id]

			if clientFind != nil {
				jsonResp, err := json.Marshal(clientFind)
				if err != nil {
					log.Printf("Error happened in JSON marshal. Err: %s\n", err)
				}
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonResp)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}

		return

	case "POST":
		id := req.URL.Query().Get("Id")
		//	fmt.Println(id)
		// clientFind := clientsList[id]
		var client client
		err := json.NewDecoder(req.Body).Decode(&client)
		// fmt.Println(client)
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
		for key := range clientsList {
			if key == id {
				delete(clientsList, id)
				clientsList[id] = append(clientsList[id], client)
				return
			}
		}
		clientsList[fmt.Sprint(len(clientsList)+1)] = append(clientsList[fmt.Sprint(len(clientsList)+1)], client)
		//	fmt.Println(clientsList)
		w.Write(jsonResp)

	case "DELETE":

		id := req.URL.Query().Get("Id")
		for key := range clientsList {
			if key == id {
				delete(clientsList, id)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte{})
		return
	}
}

func main() {

	farkim := client{Name: "Farkim", Address: "Pasaje", Phone: "123456", Email: "farkim@farkim.com"}
	rock := client{Name: "Rock&Fellers", Address: "Orono y Jujuy", Phone: "654321", Email: "rock@rock.com"}
	labac := client{Name: "Labac", Address: "Nordlink", Phone: "98765", Email: "labac@labac.com"}

	clientsList["1"] = append(clientsList["1"], farkim)
	clientsList["2"] = append(clientsList["2"], rock)
	clientsList["3"] = append(clientsList["3"], labac)

	http.HandleFunc("/client", handleClient)

	http.ListenAndServe(":8080", nil)
}

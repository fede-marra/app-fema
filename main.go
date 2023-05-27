package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// clientsMap := map[string][]client{
// 	"1":{Name: "Farkim", Address: "Pasaje", Phone: "123456", Email: "farkim@farkim.com"},
// 	"2":{Name: "Rock&Fellers", Address: "Orono y Jujuy", Phone: "654321", Email: "rock@rock.com"},
// 	"3":{Name: "Labac", Address: "Nordlink", Phone: "98765", Email: "labac@labac.com"}
// }

type client struct {
	Id      string `json:"id"`
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
	clientsList := []client{
		{Id: "1", Name: "Farkim", Address: "Pasaje", Phone: "123456", Email: "farkim@farkim.com"},
		{Id: "2", Name: "Rock&Fellers", Address: "Orono y Jujuy", Phone: "654321", Email: "rock@rock.com"},
		{Id: "3", Name: "Labac", Address: "Nordlink", Phone: "98765", Email: "labac@labac.com"},
	}
	switch req.Method {
	case "GET":
		id := req.URL.Query().Get("Id")
		if id == "" {
			// clients := []client{}

			jsonResp, err := json.Marshal(clientsList)
			if err != nil {
				log.Printf("Error happened in JSON marshal. Err: %s\n", err)
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
		} else {
			var clientFind *client
			for _, client := range clientsList {
				if client.Id == id {
					clientFind = &client
					break
				}
			}
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
		var client client
		err := json.NewDecoder(req.Body).Decode(&client)
		fmt.Println(client)
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
		clientsList = append(clientsList, client)
		fmt.Println(clientsList)
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

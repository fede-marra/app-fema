package main

import (
	"fmt"
	"time"
)

type client struct {
	name    string
	address string
	phone   string
	email   string
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

func main() {
	fmt.Println(client{}, resource{}, jobs{})
	
}

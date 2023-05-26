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

type data struct {
	router      string
	dns         string
	ip          string
	user        string
	pass        string
	addressPool string
	reservedIp  []string
}

type jobs struct {
	date       time.Time
	started    bool
	finished   bool
	clientName string
}

func main() {
	fmt.Println(client{}, data{}, jobs{})
}

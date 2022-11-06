package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/usrmaia/Go-API-CRUD/src/controller"
	"github.com/usrmaia/Go-API-CRUD/src/router"
)

func main() {
	var target string
	target = "172.17.0.2:9091"
	controller.ClientConnDial(target)
	controller.OpenDB()
	router.Router()

	fmt.Println("Server On:")
	fmt.Println("http://???.???.???.???:9090/")

	addr := ":9090"
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}

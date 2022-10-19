package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/usrmaia/GO-API-CRUD/src/controller"
	"github.com/usrmaia/GO-API-CRUD/src/router"
)

func main() {
	controller.OpenDB()
	router.Router()

	fmt.Println("Server On")

	addr := ":1357"
	var err error
	err = http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}

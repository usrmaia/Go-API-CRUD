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
	controller.InitTable()
	router.Router()

	fmt.Println("Server On:")
	fmt.Println("http://???.???.???.???:9090/")

	addr := ":9090"
	var err error
	err = http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}

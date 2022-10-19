package router

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/usrmaia/GO-API-CRUD/src/controller"
)

func routerParts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		controller.ReturnParts(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

func routerPart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	slice_url := strings.Split(r.URL.Path, "/")

	//  /part/id ou /part/id/
	if len(slice_url) == 3 && slice_url[2] != "" ||
		len(slice_url) == 4 && slice_url[2] != "" && slice_url[3] == "" {
		var id int
		var err error
		id, err = strconv.Atoi(slice_url[2])

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Method == "DELETE" {
			fmt.Println("chegou aqui", id)
			controller.DelPart(w, r, id)
			return
		} else if r.Method == "GET" {
			controller.ReturnAPart(w, r, id)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

func routerPartSimple(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		controller.AddPart(w, r)
		return
	} else if r.Method == "PUT" {
		controller.UpPart(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

func Router() {
	//http.HandleFunc("/", controller.Home)
	http.HandleFunc("/parts", routerParts)
	http.HandleFunc("/part", routerPartSimple)
	http.HandleFunc("/part/", routerPart)
}

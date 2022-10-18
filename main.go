package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// TODO - deve retornar json em minúsculo
type Part struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Value float32 `json:"value"`
}

var db *sql.DB

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func getPartsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var rows *sql.Rows
	var err error
	rows, err = db.Query(`select id, name, brand, value from Part`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var Parts []Part
	for rows.Next() {
		var part Part
		err = rows.Scan(&part.Id, &part.Name, &part.Brand, &part.Value)
		fmt.Println(part)

		if err != nil {
			continue
		}

		Parts = append(Parts, part)
	}

	err = rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(Parts)
	w.WriteHeader(http.StatusOK)
}

func addPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var new_part Part
	json.Unmarshal(data, &new_part)

	_, err = db.Exec(`
		insert into Part (name, brand, value) values
		(?, ?, ?)
	`, new_part.Name, new_part.Brand, new_part.Value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//json_encoder := json.NewEncoder(w)
	//json_encoder.Encode(new_part)
	w.WriteHeader(http.StatusCreated)
}

func getPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	slice_url := strings.Split(r.URL.Path, "/")

	if len(slice_url) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(slice_url[2])
	// TODO - Lidar com err (_) caso atoi nao consigo converter para inteiro
	// ex: /part/a

	var rows *sql.Rows
	var err error
	rows, err = db.Query(`
		select id, name, brand, value 
		from Part
		where %d = id
	`, id)

	var part Part
	err = rows.Scan(&part.Id, &part.Name, &part.Brand, &part.Value)

	if err != nil {
		log.Panicln(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = rows.Close()

	w.Header().Set("Content-Type", "application/json")
	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(part)
	w.WriteHeader(http.StatusAccepted)
}

func delPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		return
	}

	slice_url := strings.Split(r.URL.Path, "/")

	if len(slice_url) > 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(slice_url[3])
	//vendo se id existe
	row := db.QueryRow(`
		select id 
		from Part
		where id = ? 
	`, id)
	var temp_id int
	var err error
	err = row.Scan(&temp_id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err = db.Exec(`
		delete from Part
		where id = ?
	`, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func upPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	var up_part Part
	json.Unmarshal(data, &up_part)

	// Part existe
	row := db.QueryRow(`
		select id, name, brand, value
		from Part
		where id = ?
	`, up_part.Id)

	var temp_part Part
	err = row.Scan(&temp_part.Id, &temp_part.Name, &temp_part.Brand, &temp_part.Value)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//update
	_, err = db.Exec(`
		update Part 
		set name = ?, brand = ?, value = ?
	`, up_part.Name, up_part.Brand, up_part.Value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(up_part)
	w.WriteHeader(http.StatusOK)
}

func handler() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/parts", getPartsHandler)
	http.HandleFunc("/part/", getPartHandler)
	http.HandleFunc("/part/add", addPartHandler)
	http.HandleFunc("/part/del/", delPartHandler)
	http.HandleFunc("/part/up", upPartHandler)
}

func openDB() {
	var err error
	db, err = sql.Open("mysql", "root:xniC6DH rZcN84bxniC6DH rZcN84b@tcp(localhost:3306)/suzana_motorcycle_parts")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	openDB()
	handler()
	fmt.Println("Server On")
	addr := ":1357"
	log.Fatal(http.ListenAndServe(addr, nil))
}

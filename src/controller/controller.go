package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/usrmaia/GO-API-CRUD/src/model"
	"github.com/usrmaia/GO-API-CRUD/src/view"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ReturnParts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	var rows *sql.Rows

	var err error
	rows, err = model.DB.Query(`select id, name, brand, value from Part`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var Parts []model.Part
	for rows.Next() {
		var part model.Part
		err = rows.Scan(&part.Id, &part.Name, &part.Brand, &part.Value)

		if err != nil {
			continue
		}

		Parts = append(Parts, part)
	}

	err = rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	view.ResponseParts(w, Parts)
}

func AddPart(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var new_part model.Part
	json.Unmarshal(data, &new_part)

	var result sql.Result
	result, err = model.DB.Exec(`
		insert into Part (name, brand, value) values
		(?, ?, ?)
	`, new_part.Name, new_part.Brand, new_part.Value)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var id int64
	id, err = result.LastInsertId()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	new_part.Id = int(id)

	view.ResponsePart(w, new_part)
}

func ReturnAPart(w http.ResponseWriter, r *http.Request, id int) {
	var row *sql.Row
	row = model.DB.QueryRow(`
		select id, name, brand, value 
		from Part
		where id = ?
	`, id)

	var part model.Part
	var err error
	err = row.Scan(&part.Id, &part.Name, &part.Brand, &part.Value)

	if err != nil {
		log.Panicln(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	view.ResponsePart(w, part)
}

// TODO - Retornar json mostrando Part deletada
func DelPart(w http.ResponseWriter, r *http.Request, id int) {
	//id exists
	row := model.DB.QueryRow(`
		select id, name, brand, value
		from Part
		where id = ? 
	`, id)

	var err error
	var temp_part model.Part
	err = row.Scan(&temp_part.Id, &temp_part.Name, &temp_part.Brand, &temp_part.Value)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err = model.DB.Exec(`
		delete from Part
		where id = ?
	`, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	view.ResponsePart(w, temp_part)
}

func UpPart(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error
	data, err = ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var up_part model.Part
	json.Unmarshal(data, &up_part)

	// Part existe
	row := model.DB.QueryRow(`
		select id
		from Part
		where id = ?
		`, up_part.Id)

	var temp_id int
	err = row.Scan(&temp_id)

	if err != nil {
		fmt.Println(temp_id)
		fmt.Println(up_part)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//update
	_, err = model.DB.Exec(`
		update Part 
		set name = ?, brand = ?, value = ?
		where id = ?
		`, up_part.Name, up_part.Brand, up_part.Value, up_part.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	view.ResponsePart(w, up_part)
}

func OpenDB() {
	var err error
	model.DB, err = sql.Open("mysql", "root:250721@tcp(172.17.0.2:3306)/suzana_motorcycle_parts")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = model.DB.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func CreateTable() {
	var err error
	_, err = model.DB.Exec(`
		create table if not exists Part (
			id int not null auto_increment,
			name varchar(500) not null unique,
			brand varchar(50) not null,
			value float not null,
			primary key (id)
		)
	`)

	if err != nil {
		log.Fatal(err.Error())
	}

	/*
		_, err = model.DB.Exec(`insert into Part (name, brand, value) values
			("Luva para Motociclista Dedo Longo Tam. P Material Emborrachado e Couro, Branco/ Preto", "Multilaser", 47.64),
			("Capacete Moto R8 Pro Tork 56 Viseira Fume Preto/Vermelho", "Tork", 169.90),
			("Lenço de cabeça, Romacci Máscara facial Fleece máscara facial cachecol para exterior à prova de vento à prova de frio equipamento de equitação para máscara de inverno", "Romacci", 99.19)
		`)

		if err != nil {
			log.Fatal(err.Error())
		}
	*/
}

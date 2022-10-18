package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "root:xniC6DH rZcN84bxniC6DH rZcN84b@tcp(localhost:3306)/suzana_motorcycle_parts")

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec(`
		drop table if exists Part
	`)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec(`
		create table if not exists Part (
			id int not null auto_increment,
			name varchar(500) not null,
			brand varchar(50) not null,
			value float not null,
			primary key (id)
		)
	`)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Exec(`insert into Part (name, brand, value) values 
		("Luva para Motociclista Dedo Longo Tam. P Material Emborrachado e Couro, Branco/ Preto", "Multilaser", 47.64),
		("Capacete Moto R8 Pro Tork 56 Viseira Fume Preto/Vermelho", "Tork", 169.90),
		( "Lenço de cabeça, Romacci Máscara facial Fleece máscara facial cachecol para exterior à prova de vento à prova de frio equipamento de equitação para máscara de inverno", "Romacci", 99.19)
	`)

	if err != nil {
		log.Fatal(err.Error())
	}

}

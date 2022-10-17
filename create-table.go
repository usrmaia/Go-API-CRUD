//package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err_open := sql.Open("mysql", "root:xniC6DH rZcN84bxniC6DH rZcN84b@tcp(localhost:3306)/suzana_motorcycle_parts")

	if err_open != nil {
		log.Fatal(err_open.Error())
	}

	err_ping := db.Ping()

	if err_ping != nil {
		log.Fatal(err_ping.Error())
	}

	_, err_exec := db.Exec(`create table if not exists Part(
		id int not null auto_increment,
		name varchar(500),
		brand varchar(50),
		value float,
		primary key (id)
	)`)

	if err_exec != nil {
		log.Fatal(err_exec.Error())
	}

	_, err_insert := db.Exec(`insert into Part (name, brand, value) values 
		("Luva para Motociclista Dedo Longo Tam. P Material Emborrachado e Couro, Branco/ Preto", "Multilaser", 47.64),
		("Capacete Moto R8 Pro Tork 56 Viseira Fume Preto/Vermelho", "Tork", 169.90),
		( "Lenço de cabeça, Romacci Máscara facial Fleece máscara facial cachecol para exterior à prova de vento à prova de frio equipamento de equitação para máscara de inverno", "Romacci", 99.19)
	`)

	if err_insert != nil {
		log.Fatal(err_insert.Error())
	}

}

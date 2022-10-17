package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TODO - deve retornar json em minúsculo
type Part struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Value float32 `json:"value"`
}

var Parts []Part = []Part{
	Part{
		Id:    1,
		Name:  "Luva para Motociclista Dedo Longo Tam. P Material Emborrachado e Couro, Branco/ Preto",
		Brand: "Multilaser",
		Value: 47.64,
	},
	Part{
		Id:    2,
		Name:  "Capacete Moto R8 Pro Tork 56 Viseira Fume Preto/Vermelho",
		Brand: "Tork",
		Value: 169.90,
	},
	Part{
		Id:    3,
		Name:  "Lenço de cabeça, Romacci Máscara facial Fleece máscara facial cachecol para exterior à prova de vento à prova de frio equipamento de equitação para máscara de inverno",
		Brand: "Romacci",
		Value: 99.19,
	},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func getPartsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(Parts)
	w.WriteHeader(http.StatusOK)
}

func addPartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return
	}

	var new_part Part
	json.Unmarshal(data, &new_part)
	// TODO - Garantir que new_part seja um valor válido (não fazio ou com value negativo por exemplo)
	new_part.Id = len(Parts) + 1
	Parts = append(Parts, new_part)

	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(new_part)
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

	w.Header().Set("Content-Type", "application/json")

	for _, part := range Parts {
		if part.Id == id {
			json_encoder := json.NewEncoder(w)
			json_encoder.Encode(part)
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

}

func handler() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/parts", getPartsHandler)
	http.HandleFunc("/part/", getPartHandler)
	http.HandleFunc("/part/add", addPartHandler)
}

func main() {
	handler()
	fmt.Println("Server On")
	addr := ":1357"
	log.Fatal(http.ListenAndServe(addr, nil))
}
package view

import (
	"encoding/json"
	"net/http"

	"github.com/usrmaia/GO-API-CRUD/src/model"
)

func ResponseParts(w http.ResponseWriter, Parts []model.Part) {
	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(Parts)
}

func ResponsePart(w http.ResponseWriter, Part model.Part) {
	json_encoder := json.NewEncoder(w)
	json_encoder.Encode(Part)
}

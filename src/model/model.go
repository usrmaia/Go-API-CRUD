package model

import ()

type Part struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Value float32 `json:"value"`
}

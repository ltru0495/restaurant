package api

import (
	"net/http"
)

type IAPI interface {
	PostDishHandler(w http.ResponseWriter, r *http.Request)
	GetAllDishesHandler(w http.ResponseWriter, r *http.Request)
	GetDishHandler(w http.ResponseWriter, r *http.Request)
	DeleteDishHandler(w http.ResponseWriter, r *http.Request)
	PutDishHandler(w http.ResponseWriter, r *http.Request)
}

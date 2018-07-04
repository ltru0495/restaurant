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

	PostUserHandler(w http.ResponseWriter, r *http.Request)
	GetAllUsersHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
	DeleteUserHandler(w http.ResponseWriter, r *http.Request)
	PutUserHandler(w http.ResponseWriter, r *http.Request)

	PostMenuHandler(w http.ResponseWriter, r *http.Request)
	GetAllMenuHandler(w http.ResponseWriter, r *http.Request)
	GetDailyMenuHandler(w http.ResponseWriter, r *http.Request)
}

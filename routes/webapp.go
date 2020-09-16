package routes

import (
	"github.com/gorilla/mux"
	"github.com/scinna/server/services"
	"net/http"
)

func WebApp(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/", homeRoute(prv))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
}

func homeRoute(prv *services.Provider) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to scinna webapp"))
	}
}

package main

import (
	"net/http"

	"github.com/fdiaz7/go-basic-web/controllers"
	"github.com/gorilla/mux"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

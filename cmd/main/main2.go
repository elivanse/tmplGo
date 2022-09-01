package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func metodo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Mensaje metodo")

}

func main() {

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", metodo).Methods("GET")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening on 8080 ...")
	server.ListenAndServe()

}

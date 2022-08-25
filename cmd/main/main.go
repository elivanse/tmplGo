package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type sEncabezado struct {
	Titulo    string
	Contenido string
}

type sPie struct {
	Titulo    string
	Contenido string
}

type sServicio struct {
	Titulo    string
	Contenido string
	Publicado int
}

func main() {

	oPie := sPie{"P1", "ContenidoP1"}

	tServicio := template.New("oServicio")

	tHtml, err := tServicio.ParseGlob("../../pkg/utils/*.html")

	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()

	fileHeader := http.FileServer(http.Dir("../../pkg/utils"))

	mux.Handle("/", fileHeader)

	server := &http.Server{
		Addr:    ":9010",
		Handler: mux,
		// Buena Practica : reforzar los timeouts !!!!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())

	err = tHtml.ExecuteTemplate(os.Stdout, "tmplPie", oPie)
	if err != nil {
		panic(err)
	}
}

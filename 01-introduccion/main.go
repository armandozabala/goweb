package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo", r.Method)
	fmt.Fprintln(rw, "Hola Mundo con FRESH")
}

func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "La pagina no funciona", http.StatusNotFound)
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")

	fmt.Fprintf(rw, "Hola %s", name)
}

func main() {

	//Mux
	mux := http.NewServeMux()

	//Crear Router
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/saludar", Saludar)
	mux.HandleFunc("/page", PageNotFound)
	mux.HandleFunc("/error", Error)
	/*
		    http.HandleFunc("/", Hola)
			http.HandleFunc("/saludar", Saludar)
			http.HandleFunc("/page", PageNotFound)
			http.HandleFunc("/error", Error)
	*/
	//Crear un servidor

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en el PORT 3000")
	fmt.Println("Run server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())

}

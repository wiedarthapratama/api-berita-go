package Route

import (
	"berita/Controller"
	"berita/Middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigRoute() {
	router := mux.NewRouter()

	router.HandleFunc("/", Controller.Test).Methods("GET")

	router.HandleFunc("/kategori", Controller.Kategori).Methods("GET")
	router.HandleFunc("/kategori", Controller.AddKategori).Methods("POST")
	router.HandleFunc("/kategori/{id}", Controller.UpdateKategori).Methods("PATCH")
	router.HandleFunc("/kategori/{id}", Controller.DeleteKategori).Methods("DELETE")

	router.Use(Middleware.Apikey)

	log.Fatal(http.ListenAndServe(":1234", router))

	fmt.Println("Connected to port 1234")
}

package Route

import (
	"fmt"
	"log"
	"net/http"

	"../Controller"
	"../Middleware"

	"github.com/gorilla/mux"
)

func ConfigRoute() {
	router := mux.NewRouter()

	router.HandleFunc("/", Controller.Test).Methods("GET")

	router.HandleFunc("/kategori", Controller.Kategori).Methods("GET")
	router.HandleFunc("/kategori", Controller.AddKategori).Methods("POST")
	router.HandleFunc("/kategori/{id}", Controller.UpdateKategori).Methods("PATCH")
	router.HandleFunc("/kategori/{id}", Controller.DeleteKategori).Methods("DELETE")

	router.HandleFunc("/penulis", Controller.Penulis).Methods("GET")
	router.HandleFunc("/penulis", Controller.AddPenulis).Methods("POST")
	router.HandleFunc("/penulis/{id}", Controller.UpdatePenulis).Methods("PATCH")
	router.HandleFunc("/penulis/{id}", Controller.DeletePenulis).Methods("DELETE")

	router.HandleFunc("/berita", Controller.Berita).Methods("GET")
	router.HandleFunc("/berita", Controller.AddBerita).Methods("POST")
	router.HandleFunc("/berita/{id}", Controller.UpdateBerita).Methods("PATCH")

	router.Use(Middleware.Apikey)

	log.Fatal(http.ListenAndServe(":1234", router))

	fmt.Println("Connected to port 1234")
}
package Controller

import (
	"berita/Database"
	"berita/Helper"
	"berita/Model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Kategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		kategori         Model.Kategori
		arr_kategori     []Model.Kategori
		responseKategori Model.ResponseKategori
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, nama from kategori")
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&kategori.Id, &kategori.Nama); err != nil {
			Helper.LogError(err)
		} else {
			arr_kategori = append(arr_kategori, kategori)
		}
	}

	responseKategori.Status = 200
	responseKategori.Message = "Data Kategori"
	responseKategori.Data = arr_kategori

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseKategori)
}

func AddKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		Nama string `json:"nama"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	_, err = database.Exec("insert into kategori (nama) values (?)", In.Nama)
	Helper.LogError(err)

	status.Status = 200
	status.Comment = "Berhasil Disimpan"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func UpdateKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	type Input struct {
		Nama string `json:"nama"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	res, err := database.Exec("update kategori set nama = ? where id = ?", In.Nama, vars["id"])
	Helper.LogError(err)

	rowCnt, err := res.RowsAffected()
	Helper.LogError(err)

	if rowCnt != 0 {
		status.Status = 200
		status.Comment = "Berhasil Diupdate"
		w.WriteHeader(http.StatusOK)
	} else {
		status.Status = 404
		status.Comment = "data tidak tersedia"
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(status)
}

func DeleteKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	res, err := database.Exec("delete from kategori where id = ?", vars["id"])
	Helper.LogError(err)

	rowCnt, err := res.RowsAffected()
	Helper.LogError(err)

	if rowCnt != 0 {
		status.Status = 200
		status.Comment = "Berhasil Dihapus"
		w.WriteHeader(http.StatusOK)
	} else {
		status.Status = 404
		status.Comment = "data tidak tersedia"
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(status)
}

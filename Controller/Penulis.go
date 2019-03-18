package Controller

import (
	"encoding/json"
	"net/http"

	"../Database"
	"../Helper"
	"../Model"

	"github.com/gorilla/mux"
)

func Penulis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		penulis         Model.Penulis
		arr_penulis     []Model.Penulis
		responsePenulis Model.ResponsePenulis
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, nama, email, password from penulis")
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&penulis.Id, &penulis.Nama, &penulis.Email, &penulis.Password); err != nil {
			Helper.LogError(err)
		} else {
			arr_penulis = append(arr_penulis, penulis)
		}
	}

	responsePenulis.Status = 200
	responsePenulis.Message = "Data Penulis"
	responsePenulis.Data = arr_penulis

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responsePenulis)
}

func AddPenulis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	_, err = database.Exec("insert into penulis (nama, email, password) values (?,?,?)", In.Nama, In.Email, Helper.HashAndSalt([]byte(In.Password)))
	Helper.LogError(err)

	status.Status = 200
	status.Comment = "Berhasil Disimpan"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func UpdatePenulis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	type Input struct {
		Nama     string `json:"nama"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	res, err := database.Exec("update penulis set nama = ?, email = ?, password = ? where id = ?", In.Nama, In.Email, Helper.HashAndSalt([]byte(In.Password)), vars["id"])
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

func DeletePenulis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	res, err := database.Exec("delete from penulis where id = ?", vars["id"])
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

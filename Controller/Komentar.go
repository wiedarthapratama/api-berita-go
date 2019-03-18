package Controller

import (
	"encoding/json"
	"net/http"

	"../Database"
	"../Helper"
	"../Model"

	"github.com/gorilla/mux"
)

func Komentar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		komentar         Model.Komentar
		arr_komentar     []Model.Komentar
		responseKomentar Model.ResponseKomentar
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, berita_id, penulis_id, komentar, tanggal from komentar")
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&komentar.Id, &komentar.Berita_id, &komentar.Penulis_id, &komentar.Komentar, &komentar.Tanggal); err != nil {
			Helper.LogError(err)
		} else {
			arr_komentar = append(arr_komentar, komentar)
		}
	}

	responseKomentar.Status = 200
	responseKomentar.Message = "Data Komentar"
	responseKomentar.Data = arr_komentar

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseKomentar)
}

func AddKomentar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		Berita_id  string `json:"berita_id"`
		Penulis_id string `json:"penulis_id"`
		Komentar   string `json:"komentar"`
		Tanggal    string `json:"tanggal"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	_, err = database.Exec("insert into komentar (berita_id, penulis_id, komentar, tanggal) values (?,?,?,?)", In.Berita_id, In.Penulis_id, In.Komentar, In.Tanggal)
	Helper.LogError(err)

	status.Status = 200
	status.Comment = "Berhasil Disimpan"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func UpdateKomentar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	type Input struct {
		Berita_id  string `json:"berita_id"`
		Penulis_id string `json:"penulis_id"`
		Komentar   string `json:"komentar"`
		Tanggal    string `json:"tanggal"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	res, err := database.Exec("update komentar set berita_id = ?, penulis_id = ?, komentar = ?, tanggal = ? where id = ?", In.Berita_id, In.Penulis_id, In.Komentar, In.Tanggal, vars["id"])
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

func DeleteKomentar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	res, err := database.Exec("delete from komentar where id = ?", vars["id"])
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

func KomentarWhereBerita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		komentar         Model.Komentar
		arr_komentar     []Model.Komentar
		responseKomentar Model.ResponseKomentar
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, berita_id, penulis_id, komentar, tanggal from komentar where berita_id = ?", vars["id"])
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&komentar.Id, &komentar.Berita_id, &komentar.Penulis_id, &komentar.Komentar, &komentar.Tanggal); err != nil {
			Helper.LogError(err)
		} else {
			arr_komentar = append(arr_komentar, komentar)
		}
	}

	responseKomentar.Status = 200
	responseKomentar.Message = "Data Komentar"
	responseKomentar.Data = arr_komentar

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseKomentar)
}

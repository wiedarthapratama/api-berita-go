package Controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../Database"
	"../Helper"
	"../Model"
	"github.com/gorilla/mux"
)

func Suka(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		suka         Model.Suka
		arr_suka     []Model.Suka
		responseSuka Model.ResponseSuka
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, berita_id, penulis_id, jumlah_suka from suka")
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&suka.Id, &suka.Berita_id, &suka.Penulis_id, &suka.Jumlah_suka); err != nil {
			Helper.LogError(err)
		} else {
			arr_suka = append(arr_suka, suka)
		}
	}

	responseSuka.Status = 200
	responseSuka.Message = "Data Suka"
	responseSuka.Data = arr_suka

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSuka)
}

func LikeSuka(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		Berita_id  string `json:"berita_id"`
		Penulis_id string `json:"penulis_id"`
	}

	var (
		In     Input
		status Model.Status
		suka   Model.Suka
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	err = database.QueryRow("select berita_id, penulis_id, jumlah_suka from suka where berita_id = ? order by id desc limit 1", In.Berita_id).Scan(&suka.Berita_id, &suka.Penulis_id, &suka.Jumlah_suka)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if suka.Berita_id != "" {
		jumlah_suka, err := strconv.Atoi(suka.Jumlah_suka)
		Helper.LogError(err)
		hasil_suka := jumlah_suka + 1

		_, err = database.Exec("insert into suka (berita_id, penulis_id, jumlah_suka) values (?,?,?)", In.Berita_id, In.Penulis_id, hasil_suka)
		Helper.LogError(err)
	} else {
		_, err = database.Exec("insert into suka (berita_id, penulis_id, jumlah_suka) values (?,?,?)", In.Berita_id, In.Penulis_id, 1)
		Helper.LogError(err)
	}

	status.Status = 200
	status.Comment = "Berhasil Di Like"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func UnLikeSuka(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		Berita_id  string `json:"berita_id"`
		Penulis_id string `json:"penulis_id"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	res, err := database.Exec("delete from suka where berita_id = ? and penulis_id = ? ", In.Berita_id, In.Penulis_id)
	Helper.LogError(err)

	rowCnt, err := res.RowsAffected()
	Helper.LogError(err)

	if rowCnt != 0 {
		status.Status = 200
		status.Comment = "Berhasil Di UnLike"
		w.WriteHeader(http.StatusOK)
	} else {
		status.Status = 404
		status.Comment = "data tidak tersedia"
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(status)
}

func SukaWhereBerita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		suka         Model.Suka
		arr_suka     []Model.Suka
		responseSuka Model.ResponseSuka
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, berita_id, penulis_id, jumlah_suka from suka where berita_id = ?", vars["id"])
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&suka.Id, &suka.Berita_id, &suka.Penulis_id, &suka.Jumlah_suka); err != nil {
			Helper.LogError(err)
		} else {
			arr_suka = append(arr_suka, suka)
		}
	}

	responseSuka.Status = 200
	responseSuka.Message = "Data Suka where berita"
	responseSuka.Data = arr_suka

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSuka)
}

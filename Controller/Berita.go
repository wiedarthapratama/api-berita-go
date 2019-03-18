package Controller

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"../Database"
	"../Helper"
	"../Model"
	"github.com/gorilla/mux"
	"github.com/vjeantet/jodaTime"
)

func Berita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		berita         Model.Berita
		arr_berita     []Model.Berita
		responseberita Model.ResponseBerita
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	rows, err := database.Query("select id, judul, foto from berita")
	Helper.LogError(err)

	for rows.Next() {
		if err := rows.Scan(&berita.Id, &berita.Judul, &berita.Foto); err != nil {
			Helper.LogError(err)
		} else {
			arr_berita = append(arr_berita, berita)
		}
	}

	responseberita.Status = 200
	responseberita.Message = "Data berita"
	responseberita.Data = arr_berita

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseberita)
}

func AddBerita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Input struct {
		Judul       string `json:"judul"`
		Isi         string `json:"isi"`
		Filename    string `json:"filename"`
		Kategori_id string `json:"kategori_id"`
		Penulis_id  string `json:"penulis_id"`
		Foto        string `json:"foto"`
	}

	var (
		In     Input
		status Model.Status
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	date := jodaTime.Format("YYYYMMddHHmmss", time.Now())
	filename := In.Penulis_id + "-" + date
	data, err := base64.StdEncoding.DecodeString(In.Foto)
	filename = filename + In.Filename[len(In.Filename)-4:]

	err = ioutil.WriteFile(Helper.FolderPath+filename, data, 0644)

	_, err = database.Exec("insert into Berita (judul, isi, foto, kategori_id, penulis_id) values (?,?,?,?,?)", In.Judul, In.Isi, filename, In.Kategori_id, In.Penulis_id)
	Helper.LogError(err)

	status.Status = 200
	status.Comment = "Berhasil Disimpan"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func UpdateBerita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	type Input struct {
		Judul       string `json:"judul"`
		Isi         string `json:"isi"`
		Filename    string `json:"filename"`
		Kategori_id string `json:"kategori_id"`
		Penulis_id  string `json:"penulis_id"`
		Foto        string `json:"foto"`
	}

	var (
		In     Input
		status Model.Status
		berita Model.Berita
		res    sql.Result
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := json.NewDecoder(r.Body).Decode(&In)
	Helper.LogError(err)

	if In.Filename != "" {
		err := database.QueryRow("select foto from berita where id = ?", vars["id"]).Scan(&berita.Foto)
		Helper.LogError(err)
		os.Remove(Helper.FolderPath + berita.Foto)

		date := jodaTime.Format("YYYYMMddHHmmss", time.Now())
		filename := In.Penulis_id + "-" + date
		data, err := base64.StdEncoding.DecodeString(In.Foto)
		filename = filename + In.Filename[len(In.Filename)-4:]

		err = ioutil.WriteFile(Helper.FolderPath+filename, data, 0644)

		res, err = database.Exec("update berita set judul = ?, isi = ?, foto = ?, kategori_id = ?, penulis_id = ? where id = ?", In.Judul, In.Isi, filename, In.Kategori_id, In.Penulis_id, vars["id"])
		Helper.LogError(err)
	} else {
		res, err = database.Exec("update berita set judul = ?, isi = ?, kategori_id = ?, penulis_id = ? where id = ?", In.Judul, In.Isi, In.Kategori_id, In.Penulis_id, vars["id"])
		Helper.LogError(err)
	}

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

func DeleteBerita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		status Model.Status
		berita Model.Berita
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := database.QueryRow("select foto from berita where id = ?", vars["id"]).Scan(&berita.Foto)
	Helper.LogError(err)
	os.Remove(Helper.FolderPath + berita.Foto)

	res, err := database.Exec("delete from berita where id = ?", vars["id"])
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

func ReadBerita(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	var (
		responseReadBerita Model.ResponseReadBerita
		readBerita         Model.ReadBerita
	)

	database := Database.ConfigDatabase()
	defer database.Close()

	err := database.QueryRow("select id, judul, isi, foto, terbaca, kategori_id, penulis_id from berita where id = ?", vars["id"]).Scan(&readBerita.Id, &readBerita.Judul, &readBerita.Isi, &readBerita.Foto, &readBerita.Terbaca, &readBerita.Kategori_id, &readBerita.Penulis_id)
	Helper.LogError(err)

	responseReadBerita.Status = 200
	responseReadBerita.Message = "Berhasil Ditampilkan"
	responseReadBerita.Data = readBerita
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(responseReadBerita)
}

package Model

type Komentar struct {
	Id         int    `json:"id"`
	Berita_id  string `json:"berita_id"`
	Penulis_id string `json:"penulis_id"`
	Komentar   string `json:"komentar"`
	Tanggal    string `json:"tanggal"`
}

type ResponseKomentar struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Komentar
}

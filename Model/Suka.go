package Model

type Suka struct {
	Id          int    `json:"id"`
	Berita_id   string `json:"berita_id"`
	Penulis_id  string `json:"penulis_id"`
	Jumlah_suka string `json:"jumlah_suka"`
}

type ResponseSuka struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Suka
}

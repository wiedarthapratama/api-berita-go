package Model

type Berita struct {
	Id    int    `json:"id"`
	Judul string `json:"judul"`
	Foto  string `json:"foto"`
}

type ResponseBerita struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Berita
}

type ReadBerita struct {
	Id          int    `json:"id"`
	Judul       string `json:"judul"`
	Isi         string `json:"isi"`
	Foto        string `json:"foto"`
	Terbaca     string `json:"terbaca"`
	Kategori_id string `json:"kategori_id"`
	Penulis_id  string `json:"penulis_id"`
}

type ResponseReadBerita struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    ReadBerita
}

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

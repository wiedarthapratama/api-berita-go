package Model

type Kategori struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type ResponseKategori struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Kategori
}

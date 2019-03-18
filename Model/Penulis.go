package Model

type Penulis struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponsePenulis struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Penulis
}

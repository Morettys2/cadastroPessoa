package models

type Pessoa struct {
	ID            string `json:"id"`
	Nome          string `json:"nome"`
	Idade         int    `json:"idade"`
	Sexo          string `json:"sexo"`
	CoisaFavorita string `json:"coisa_favorita"`
}

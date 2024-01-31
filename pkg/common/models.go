package common

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Nome           string `json:"nome"`
	Genero         string `json:"genero"`
	Nascimento     string `json:"aniversario"`
	CPF            string `json:"cpf"`
	Telefone       string `json:"telefone"`
	Endereco       string `json:"endereco"`
	NumeroEndereco string `json:"numero"`
	Estado         string `json:"estado"`
}

type Chamado struct {
	gorm.Model
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
}

type Cliente struct {
	gorm.Model
	Nome           string `json:"nome"`
	Genero         string `json:"genero"`
	Nascimento     string `json:"aniversario"`
	CPF            string `json:"cpf"`
	Telefone       string `json:"telefone"`
	Endereco       string `json:"endereco"`
	NumeroEndereco string `json:"numero"`
	Estado         string `json:"estado"`
}

type Menu struct {
	Id       int           `json:"id"`
	Label    string        `json:"label"`
	Link     string        `json:"link"`
	Icon     string        `json:"icon"`
	SubItems []SubMenuItem `json:"subItems"`
}

type SubMenuItem struct {
	Label string `json:"label"`
	Link  string `json:"link"`
	Icon  string `json:"icon"`
}

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

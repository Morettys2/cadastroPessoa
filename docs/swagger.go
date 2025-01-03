package docs

import "github.com/swaggo/swag"

// swagger:route GET /swagger/index.html docs getSwagger
// Swagger documentation for the API

var doc = `{
	"swagger": "2.0",
	"info": {
		"description": "Esta API permite cadastrar, editar e deletar pessoas.",
		"title": "API de Cadastro de Pessoas",
		"version": "1.0"
	},
	"host": "localhost:8080",
	"basePath": "/"
}`

// CustomSwagger é uma implementação personalizada da interface Swagger
type CustomSwagger struct{}

// ReadDoc retorna a documentação Swagger
func (c CustomSwagger) ReadDoc() string {
	return doc
}

func init() {
	swag.Register(swag.Name, &CustomSwagger{})
}

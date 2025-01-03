package main

import (
	"project/pkg/router"
	"project/internal/db"
	_ "project/docs" // Isso vai carregar automaticamente a documentação Swagger
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Inicializa o banco de dados
	db.InitDB()

	// Configura as rotas do servidor
	r := router.Setup()

	// Serve a documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia o servidor
	r.Run(":8080")
}

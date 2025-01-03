package router

import (
	"project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/pessoas", handlers.ListarPessoas)
	r.POST("/pessoas", handlers.CriarPessoa)
	r.PUT("/pessoas/:id", handlers.AtualizarPessoa)
	r.DELETE("/pessoas/:id", handlers.DeletarPessoa)

	return r
}

package handlers

import (
	"net/http"
	"project/internal/db"
	"project/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListarPessoas(c *gin.Context) {
	rows, err := db.DB.Query("SELECT * FROM pessoas")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var pessoas []models.Pessoa
	for rows.Next() {
		var p models.Pessoa
		rows.Scan(&p.ID, &p.Nome, &p.Idade, &p.Sexo, &p.CoisaFavorita)
		pessoas = append(pessoas, p)
	}
	c.JSON(http.StatusOK, pessoas)
}


func CriarPessoa(c *gin.Context) {
	var novaPessoa models.Pessoa
	if err := c.ShouldBindJSON(&novaPessoa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	novaPessoa.ID = uuid.New().String()
	_, err := db.DB.Exec("INSERT INTO pessoas (id, nome, idade, sexo, coisa_favorita) VALUES (?, ?, ?, ?, ?)", novaPessoa.ID, novaPessoa.Nome, novaPessoa.Idade, novaPessoa.Sexo, novaPessoa.CoisaFavorita)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, novaPessoa)
}


func AtualizarPessoa(c *gin.Context) {
	id := c.Param("id")
	var pessoaAtualizada models.Pessoa
	if err := c.ShouldBindJSON(&pessoaAtualizada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec("UPDATE pessoas SET nome=?, idade=?, sexo=?, coisa_favorita=? WHERE id=?", pessoaAtualizada.Nome, pessoaAtualizada.Idade, pessoaAtualizada.Sexo, pessoaAtualizada.CoisaFavorita, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pessoaAtualizada.ID = id
	c.JSON(http.StatusOK, pessoaAtualizada)
}

func DeletarPessoa(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM pessoas WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

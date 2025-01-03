package handlers

import (
	"net/http"
	"project/internal/db"
	"project/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListarPessoas godoc
// @Summary Lista todas as pessoas
// @Description Retorna uma lista de todas as pessoas cadastradas
// @Tags pessoas
// @Produce json
// @Success 200 {array} models.Pessoa
// @Router /pessoas [get]
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

// CriarPessoa godoc
// @Summary Cria uma nova pessoa
// @Description Cadastra uma nova pessoa no banco de dados
// @Tags pessoas
// @Accept json
// @Produce json
// @Param pessoa body models.Pessoa true "Pessoa para cadastrar"
// @Success 201 {object} models.Pessoa
// @Router /pessoas [post]
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

// AtualizarPessoa godoc
// @Summary Atualiza uma pessoa existente
// @Description Atualiza os dados de uma pessoa pelo ID
// @Tags pessoas
// @Accept json
// @Produce json
// @Param id path string true "ID da Pessoa"
// @Param pessoa body models.Pessoa true "Dados atualizados"
// @Success 200 {object} models.Pessoa
// @Router /pessoas/{id} [put]
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

// DeletarPessoa godoc
// @Summary Deleta uma pessoa
// @Description Remove uma pessoa do banco de dados pelo ID
// @Tags pessoas
// @Param id path string true "ID da Pessoa"
// @Success 204 "No Content"
// @Router /pessoas/{id} [delete]
func DeletarPessoa(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM pessoas WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

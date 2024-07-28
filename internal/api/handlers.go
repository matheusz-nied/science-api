package api

import (
	"net/http"
	"nied-science/internal/model"
	"nied-science/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAPODs(c *gin.Context) {
	// Obter parâmetros de consulta
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// Buscar APODs com paginação
	repo := repository.NewAPODRepository()

	apods, totalItems, err := repo.GetAPODsPaginated(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calcular o número total de páginas
	totalPages := int(totalItems / int64(pageSize))
	if totalItems%int64(pageSize) != 0 {
		totalPages++
	}

	// Criar a resposta paginada
	response := model.PaginatedResponse{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Items:      apods,
	}

	c.JSON(http.StatusOK, response)
}

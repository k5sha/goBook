package books

import (
	"github.com/k5sha/goBook/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &book)
}

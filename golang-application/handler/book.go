package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/shatvik/week2-GL1-CipherSchools.git/database"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB)
	if err != nil {
		log.Fatal(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k5sha/goBook/internal/books"
	"github.com/k5sha/goBook/internal/config"
	"github.com/k5sha/goBook/internal/storage"
)

func main() {

	port := config.Get().Port
	dbUrl := config.Get().DatabaseDSN

	r := gin.Default()
	h := storage.Init(dbUrl)

	books.RegisterRoutes(r, h)

	r.Run(port)
}

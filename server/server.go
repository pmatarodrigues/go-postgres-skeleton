package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmatarodrigues/go-starter/storage"
)

func New(repository *storage.Repo) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", getStatus)
	return r
}

func getStatus(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

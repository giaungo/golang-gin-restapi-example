package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func main() {
	indexComments()

	router := gin.Default()
	router.GET("/topPosts", getTopPosts)
	router.GET("/search", search)
	router.Run("localhost:8080")
}

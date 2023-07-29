package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/q10357/RelService/graph"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/rel", graph.RelGraphRouter)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	router.Run("127.0.0.1:" + port)

}

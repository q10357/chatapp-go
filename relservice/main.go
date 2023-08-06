package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/q10357/RelService/data/rel"
	"github.com/q10357/RelService/data/user"
	"github.com/q10357/RelService/services"
	"github.com/q10357/RelService/web/graph"
	"github.com/q10357/RelService/web/middleware"
	"github.com/q10357/RelService/web/schemas"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	// create your schema here
	relRepo := rel.NewRelRepo()
	userRepo := user.NewUserRepo()

	relService := services.NewRelService(relRepo, userRepo)
	relSchema, err := schemas.NewRelRootSchema(relService)

	if err != nil {
		log.Fatalf("Error creating rel schema %d", err)
	}

	//router.Use(cors.Default())
	router.Use(middleware.ValidateHeaders())
	router.POST("/rel", graph.NewRelGraphRouter(relSchema))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	router.Run(":" + port)

}

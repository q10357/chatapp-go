package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/q10357/RelService/data/rel"
	"github.com/q10357/RelService/data/user"
	"github.com/q10357/RelService/services"
	"github.com/q10357/RelService/web/graph"
	"github.com/q10357/RelService/web/middleware"
	"github.com/q10357/RelService/web/schemas"
)

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	relRepo := rel.NewRelRepo()
	userRepo := user.NewUserRepo()

	relService := services.NewRelService(relRepo, userRepo)
	relSchema, err := schemas.NewRelRootSchema(relService)
	if err != nil {
		log.Fatalf("Error creating rel schema: %v", err)
	}

	router.Use(middleware.ValidateHeaders())
	router.POST("/rel", graph.NewRelGraphRouter(relSchema))

	isDevEnv, err := strconv.ParseBool(os.Getenv("DEV_ENV"))
	if err != nil {
		log.Fatalf("Error converting env variable: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	host := "127.0.0.1"
	if !isDevEnv {
		host = ""
	}

	router.Run(fmt.Sprintf("%s:%s", host, port))
}

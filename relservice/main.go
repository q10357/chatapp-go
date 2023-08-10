package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/q10357/RelService/database"
	"github.com/q10357/RelService/database/data/rel"
	"github.com/q10357/RelService/database/data/user"
	"github.com/q10357/RelService/services"
	"github.com/q10357/RelService/web/graph"
	"github.com/q10357/RelService/web/middleware"
	"github.com/q10357/RelService/web/schemas"
)

type Config struct {
	DBUser string
	DBPass string
	DBName string
	PORT   string
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	dbCfg := &database.Config{
		DBUser: cfg.DBUser,
		DBPass: cfg.DBPass,
		DBName: cfg.DBName,
	}

	db, err := database.InitDatabase(dbCfg)
	if err != nil {
		log.Fatal(err)
	}

	router := setupRouter(db, cfg)
	fmt.Printf("listening on port %s\n", cfg.PORT)
	router.Run(fmt.Sprintf(":%s", cfg.PORT))
}

func loadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8082"
	}

	cfg := &Config{
		DBUser: os.Getenv("DBUSER"),
		DBPass: os.Getenv("DBPASS"),
		DBName: os.Getenv("DBNAME"),
		PORT: os.Getenv("PORT"),
	}

	return cfg, nil
}

func setupRouter(db *sql.DB, cfg *Config) *gin.Engine {
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

	return router
}

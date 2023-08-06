package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

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
	DevEnv bool
	Port   string
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

	host := "127.0.0.1"
	if !cfg.DevEnv {
		host = ""
	}

	router.Run(fmt.Sprintf("%s:%s", host, cfg.Port))
}

func loadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	isDevEnv, err := strconv.ParseBool(os.Getenv("DEV_ENV"))
	if err != nil {
		log.Fatalf("Error converting env variable: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	cfg := &Config{
		DBUser: os.Getenv("DBUSER"),
		DBPass: os.Getenv("DBPASS"),
		DBName: os.Getenv("DBNAME"),
		DevEnv: isDevEnv,
		Port:   port,
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

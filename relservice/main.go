package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

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

	db, err := setupDatabase(cfg)
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

func setupDatabase(cfg *Config) (*sql.DB, error) {
	dbCfg := mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPass,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: cfg.DBName,
	}

	db, err := sql.Open("mysql", dbCfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return db, nil
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

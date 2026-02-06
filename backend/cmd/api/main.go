package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/josediaz/go-and-reactjs/backend/internal/application"
	"github.com/josediaz/go-and-reactjs/backend/internal/domain"
	httpAdapter "github.com/josediaz/go-and-reactjs/backend/internal/infrastructure/adapter/http"
	"github.com/josediaz/go-and-reactjs/backend/internal/infrastructure/adapter/repository"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/demo?sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("database ping:", err)
	}

	// Capa de dominio -> aplicación -> adaptadores (Hexagonal)
	var repo domain.TaskRepository = repository.NewPostgresTaskRepository(db)
	svc := application.NewTaskService(repo)
	handler := httpAdapter.NewTaskHandler(svc)

	// GraphQL schema
	schema, err := httpAdapter.BuildTaskGraphQLSchema(svc)
	if err != nil {
		log.Fatal("graphql schema:", err)
	}

	r := gin.Default()
	r.Use(corsMiddleware())

	// REST API
	api := r.Group("/api/v1/tasks")
	{
		api.GET("", handler.List)
		api.GET("/:id", handler.GetByID)
		api.POST("", handler.Create)
		api.PATCH("/:id/status", handler.UpdateStatus)
		api.DELETE("/:id", handler.Delete)
	}

	// GraphQL endpoint
	r.POST("/graphql", gin.WrapF(func(w http.ResponseWriter, req *http.Request) {
		var payload struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  payload.Query,
			VariableValues: payload.Variables,
		})
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(result)
	}))

	// Health para Kubernetes
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("API listening on :" + port)
	log.Fatal(r.Run(":" + port))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

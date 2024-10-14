package api

import (
	"database/sql"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luytbq/go-angular-reactive-crud-example/pkg/category"
)

func (server *APIServer) Run() error {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	categoryHandler := category.NewCategoryHandler(server.DB)

	categoryHandler.RegisterRoute(engine, server.Prefix)

	err := engine.Run(server.Port)

	return err
}

func NewAPIServer(port string, prefix string, db *sql.DB) *APIServer {
	return &APIServer{
		Port:   port,
		Prefix: prefix,
		DB:     db,
	}
}

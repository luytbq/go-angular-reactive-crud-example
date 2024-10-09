package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/luytbq/go-angular-reactive-crud-example/pkg/category"
)

func (server *APIServer) Run() error {
	engine := gin.Default()

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

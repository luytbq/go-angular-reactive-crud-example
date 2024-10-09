package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Port   string
	Prefix string
	DB     *sql.DB
}

type BaseEntity[IDType any] interface {
	GetID() IDType
	// SetID(id IDType)
}

type Handler interface {
	RegisterRoute(engine *gin.Engine, prefix string)
}

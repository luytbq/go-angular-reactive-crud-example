package category

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	Repo Repository
}

func NewCategoryHandler(db *sql.DB) *CategoryHandler {
	return &CategoryHandler{
		Repo: NewRepositoryImpl(db),
	}
}

func (handler CategoryHandler) RegisterRoute(engine *gin.Engine, prefix string) {
	engine.GET(prefix+"/categories/:id", handler.handleGetByID)
	engine.GET(prefix+"/categories", handler.handleQuery)

}

func (handler CategoryHandler) handleGetByID(context *gin.Context) {
	var id uint64
	id, err := strconv.ParseUint(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := handler.Repo.getById(id)
	if err != nil {
		log.Printf("category handleGetById error %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	context.JSON(http.StatusOK, category)
}

func (handler CategoryHandler) handleQuery(context *gin.Context) {
	var params CategoryQueryParams
	params.Keyword = context.Query("keyword")
	params.Page, _ = strconv.Atoi(context.DefaultQuery("page", "1"))
	params.PageSize, _ = strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	result, err := handler.Repo.query(&params)

	if err != nil {
		log.Printf("category handleQuery error %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, result)
}

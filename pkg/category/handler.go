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
	engine.GET(prefix+"/categories", handler.handleSearch)
	engine.POST(prefix+"/categories", handler.handleCreate)
	engine.PATCH(prefix+"/categories", handler.handleUpdate)

}

func (handler CategoryHandler) handleUpdate(context *gin.Context) {
	var category Category
	if err := context.BindJSON(&category); err != nil {
		log.Printf("category handleUpdate error: %v", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := handler.Repo.update(&category); err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		} else if err == ErrResourceExisted {
			context.JSON(http.StatusBadRequest, gin.H{"error": "category existed"})
			return
		}
		log.Printf("category handleUpdate error: %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, category)
}

func (handler CategoryHandler) handleCreate(context *gin.Context) {
	var category Category
	if err := context.BindJSON(&category); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if err := handler.Repo.create(&category); err != nil {
		log.Printf("category handleCreate error: %v", err)
		if err == ErrResourceExisted {
			context.JSON(http.StatusBadRequest, gin.H{"error": "category existed"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusCreated, category)
}

func (handler CategoryHandler) handleGetByID(context *gin.Context) {
	var id uint64
	id, err := strconv.ParseUint(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	category, err := handler.Repo.searchById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		log.Printf("category handleGetById error %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, category)
}

func (handler CategoryHandler) handleSearch(context *gin.Context) {
	var params CategorySearchParams
	params.Keyword = context.Query("keyword")
	params.Page, _ = strconv.Atoi(context.DefaultQuery("page", "1"))
	params.PageSize, _ = strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	result, err := handler.Repo.search(&params)

	if err != nil {
		log.Printf("category handleSearch error %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, result)
}

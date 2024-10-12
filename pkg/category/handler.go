package category

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/luytbq/go-angular-reactive-crud-example/pkg/common"

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
	engine.POST(prefix+"/categories", handler.handleCreate)
	engine.GET(prefix+"/categories/:id", handler.handleGetByID)
	engine.GET(prefix+"/categories", handler.handleSearch)
	engine.PATCH(prefix+"/categories", handler.handleUpdate)
	engine.DELETE(prefix+"/categories/:id", handler.handleDelete)
}

func (handler CategoryHandler) handleDelete(context *gin.Context) {
	var id uint64
	id, err := strconv.ParseUint(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, common.ResponseInvalidID)
		return
	}

	err = handler.Repo.delete(id)
	if err != nil {
		if err == common.ErrResourceNotFound {
			context.JSON(http.StatusNotFound, common.ResponseResourceNotFound)
			return
		}
		log.Printf("category handleDelete error: %v", err)
		context.JSON(http.StatusInternalServerError, common.ResponseInternalError)
		return
	}
	context.Status(200)
}

func (handler CategoryHandler) handleUpdate(context *gin.Context) {
	var category Category
	if err := context.BindJSON(&category); err != nil {
		log.Printf("category handleUpdate error: %v", err)
		context.JSON(http.StatusBadRequest, common.ResponseBadRequest)
		return
	}

	if err := validateCategory(&category); err != nil {
		context.JSON(http.StatusBadRequest, common.ResponseError(err))
		return
	}

	if err := handler.Repo.update(&category); err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, common.ResponseResourceNotFound)
			return
		} else if err == common.ErrResourceExisted {
			context.JSON(http.StatusBadRequest, common.ResponseResourceExisted)
			return
		}
		log.Printf("category handleUpdate error: %v", err)
		context.JSON(http.StatusInternalServerError, common.ResponseInternalError)
		return
	}

	context.JSON(http.StatusOK, category)
}

func (handler CategoryHandler) handleCreate(context *gin.Context) {
	var category Category
	if err := context.BindJSON(&category); err != nil {
		context.JSON(http.StatusBadRequest, common.ResponseBadRequest)
		return
	}

	if err := validateCategory(&category); err != nil {
		context.JSON(http.StatusBadRequest, common.ResponseError(err))
		return
	}

	if err := handler.Repo.create(&category); err != nil {
		log.Printf("category handleCreate error: %v", err)
		if err == common.ErrResourceExisted {
			context.JSON(http.StatusBadRequest, common.ResponseResourceExisted)
			return
		}
		context.JSON(http.StatusInternalServerError, common.ResponseInternalError)
		return
	}

	context.JSON(http.StatusCreated, category)
}

func (handler CategoryHandler) handleGetByID(context *gin.Context) {
	var id uint64
	id, err := strconv.ParseUint(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, common.ResponseInvalidID)
		return
	}

	category, err := handler.Repo.searchById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, common.ResponseResourceNotFound)
			return
		}
		log.Printf("category handleGetById error %v", err)
		context.JSON(http.StatusInternalServerError, common.ResponseInternalError)
		return
	}

	context.JSON(http.StatusOK, category)
}

func (handler CategoryHandler) handleSearch(context *gin.Context) {
	var params CategorySearchParams
	params.Keyword = context.Query("keyword")

	pageSize, err := strconv.Atoi(context.DefaultQuery("pageSize", "-1"))
	if err != nil {
		pageSize = 10
	}
	if params.PageSize < 1 {
		params.PageSize = -1
	}
	params.PageSize = pageSize

	page, err := strconv.Atoi(context.DefaultQuery("page", "1"))
	if err != nil {
		page = 10
	}
	params.Page = page

	result, err := handler.Repo.search(&params)

	if err != nil {
		if err == common.ErrResourceNotFound {
			context.JSON(http.StatusNotFound, common.ResponseResourceNotFound)
			return
		}
		log.Printf("category handleSearch error %v", err)
		context.JSON(http.StatusInternalServerError, common.ResponseInternalError)
		return
	}

	context.JSON(http.StatusOK, result)
}

func validateCategory(category *Category) error {
	if category.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}
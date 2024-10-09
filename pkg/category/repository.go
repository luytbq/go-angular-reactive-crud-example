package category

import (
	"database/sql"
)

type RepositoryImpl struct {
	DB *sql.DB
}

func NewRepositoryImpl(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		DB: db,
	}
}

func (repo RepositoryImpl) getById(id uint64) (*Category, error) {
	return &Category{}, nil
}

func (repo RepositoryImpl) create(category *Category) (*Category, error) {
	return &Category{}, nil
}

func (repo RepositoryImpl) update(category *Category) (*Category, error) {

	return &Category{}, nil
}

func (repo RepositoryImpl) query(params *CategoryQueryParams) (*CategoryQueryResponse, error) {
	// return nil, errors.New("fake error")
	return &CategoryQueryResponse{}, nil
}

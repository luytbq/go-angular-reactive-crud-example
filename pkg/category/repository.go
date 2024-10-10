package category

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrResourceExisted = errors.New("resource existed")
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

func (repo RepositoryImpl) create(category *Category) error {
	tx, err := repo.DB.BeginTx(context.TODO(), &sql.TxOptions{})
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return err
	}

	querySelect := `select id from categories where name = $1`
	_, err = tx.Query(querySelect, category.Name)

	// sql.ErrNoRows is expected
	if err == nil {
		return ErrResourceExisted
	} else if err != sql.ErrNoRows {
		return err
	}

	queryInsert := `insert into categories(name) values($1) returning id`
	err = repo.DB.QueryRow(queryInsert, category.Name).Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo RepositoryImpl) update(category *Category) (*Category, error) {

	return &Category{}, nil
}

func (repo RepositoryImpl) query(params *CategoryQueryParams) (*CategoryQueryResponse, error) {
	// return nil, errors.New("fake error")
	return &CategoryQueryResponse{}, nil
}

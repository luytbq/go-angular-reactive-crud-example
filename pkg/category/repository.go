package category

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/luytbq/go-angular-reactive-crud-example/pkg/common"
)

var (
	sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

type RepositoryImpl struct {
	DB *sql.DB
}

func NewRepositoryImpl(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		DB: db,
	}
}

func (repo RepositoryImpl) searchById(id uint64) (*Category, error) {
	category := &Category{ID: id}
	querySelect := `select name from categories where id = $1`

	if err := repo.DB.QueryRow(querySelect, category.ID).Scan(&category.Name); err != nil {
		return nil, err
	}

	return category, nil
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

	// first, check if name existed
	querySelect := `select id from categories where name = $1`
	rows, err := tx.Query(querySelect, category.Name)

	// sql.ErrNoRows is expected
	if err == nil {
		for rows.Next() {
			return common.ErrResourceExisted
		}
	} else if err != sql.ErrNoRows {
		return err
	}

	// then execute insert
	queryInsert := `insert into categories(name) values($1) returning id`
	if err = repo.DB.QueryRow(queryInsert, category.Name).Scan(&category.ID); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo RepositoryImpl) update(category *Category) error {
	tx, err := repo.DB.BeginTx(context.TODO(), &sql.TxOptions{})
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// first, check id must existed and name must not existed
	var count int64
	querySelectById := `select count(id) from categories where id = $1 or (id != $1 and name = $2)`
	if err = tx.QueryRow(querySelectById, category.ID, category.Name).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	} else if count > 1 {
		return common.ErrResourceExisted
	}

	log.Printf("count = %d", count)

	// then execute update
	queryUpdate := `update categories
		set name = $1
		where id = $2`

	result, err := tx.Exec(queryUpdate, category.Name, category.ID)
	if err != nil {
		log.Println("aaa")
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	count, err = result.RowsAffected()
	log.Printf("rowAffected: %d", count)
	if count < 1 {
		err = sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

func (repo RepositoryImpl) delete(id uint64) error {
	query := `delete from categories where id = $1`

	result, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil || count > 1 {
		return errors.New("some thing went wrong")
	}

	if count == 0 {
		return common.ErrResourceNotFound
	}

	return nil
}

func (repo RepositoryImpl) search(params *CategorySearchParams) (*CategorySearchResponse, error) {
	log.Printf("category search params %+v", params)
	qb := sq.Select("id, name").From("categories")

	if params.PageSize > 0 {
		qb = qb.Limit(uint64(params.PageSize)).
			Offset(uint64((params.Page - 1) * params.PageSize))
	}

	if params.Keyword != "" {
		qb = qb.Where(squirrel.Like{"name": "%" + params.Keyword + "%"})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		log.Println("category search error building query")
		return nil, err
	}
	log.Printf("category search query: %s", query)
	log.Printf("category search args: %v", args)

	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		log.Println("category search error executing query")
		return nil, err
	}

	response := &CategorySearchResponse{TotalPages: 1, CurrentPage: 1, PageSize: -1}
	response.Items = make([]*Category, 0, 10)

	for rows.Next() {
		category := &Category{}
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			log.Println("category search error scanning rows")
			return nil, err
		}

		response.Items = append(response.Items, category)
	}

	if params.PageSize > 1 {
		repo.addSearchTotal(response, params)
	}

	return response, nil
}

func (repo *RepositoryImpl) addSearchTotal(response *CategorySearchResponse, params *CategorySearchParams) error {
	qb := sq.Select("count(1)").From("categories")

	if params.Keyword != "" {
		qb = qb.Where(squirrel.Like{"name": "%" + params.Keyword + "%"})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		log.Println("category search error building query")
		return err
	}
	log.Printf("category search query: %s", query)
	log.Printf("category search args: %v", args)

	var total int64 = 0
	err = repo.DB.QueryRow(query, args...).Scan(&total)
	if err != nil {
		log.Println("category search error count total")
		return err
	}

	response.TotalPages = int(total/int64(params.PageSize) + 1)
	response.CurrentPage = params.Page
	response.PageSize = params.PageSize

	return nil
}

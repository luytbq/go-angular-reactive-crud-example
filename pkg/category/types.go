package category

type Category struct {
	ID   uint64
	Name string
	// Description string
}

type CategoryQueryParams struct {
	Keyword  string
	Page     int
	PageSize int
}

type CategoryQueryResponse struct {
	CurrentPage int
	PageSize    int
	TotalPage   int
	Items       []Category
}

type Repository interface {
	getById(id uint64) (*Category, error)
	create(category *Category) error
	update(category *Category) (*Category, error)
	query(params *CategoryQueryParams) (*CategoryQueryResponse, error)
}

package category

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	// Description string
}

type CategorySearchParams struct {
	Keyword  string
	Page     int
	PageSize int
}

type CategorySearchResponse struct {
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	TotalPage   int         `json:"totalPage"`
	Items       []*Category `json:"items"`
}

type Repository interface {
	searchById(id uint64) (*Category, error)
	create(category *Category) error
	update(category *Category) error
	search(params *CategorySearchParams) (*CategorySearchResponse, error)
}

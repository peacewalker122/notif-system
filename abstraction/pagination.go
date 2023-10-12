package abstraction

import "fmt"

type Pagination struct {
	Page    int     `json:"page"`
	PerPage int     `json:"per_page"`
	OrderBy *string `json:"order_by"`
	Order   *string `json:"order" enum:"asc,desc"`
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Pagination) GetPerPage() int {
	if p.PerPage == 0 {
		return 10
	}
	return p.PerPage
}

func (p *Pagination) GetOrder() string {
	order := "asc"
	if p.Order != nil {
		order = *p.Order
	}

	return fmt.Sprintf("%s %s", *p.OrderBy, order)
}

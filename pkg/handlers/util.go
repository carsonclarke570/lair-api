package handlers

import (
	"strconv"
)

// Pagination defines organizatinal variabales for pagination
type Pagination struct {
	Order   string
	OrderBy string
	Page    uint
	N       uint
}

// NewPagination creates a default pagination object
func NewPagination() *Pagination {
	return &Pagination{
		Order:   "ASC",
		OrderBy: "id",
		Page:    1,
		N:       12,
	}
}

func (p *Pagination) parseQuery(key, value string) (bool, error) {
	if key == "order" {
		if value == "des" {
			p.Order = "DESC"
		}
		return true, nil
	} else if key == "order_by" {
		p.OrderBy = value
		return true, nil
	} else if key == "n" {
		n, err := strconv.Atoi(value)
		if err != nil {
			return false, err
		}
		p.N = uint(n)
		return true, nil
	} else if key == "page" {
		page, err := strconv.Atoi(value)
		if err != nil {
			return false, err
		}
		p.Page = uint(page)
		return true, nil
	}
	return false, nil
}

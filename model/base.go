package model

import "github.com/lingfohn/lime/ent"

type paginationQuery struct {
	Current  int `form:"current" json:"current,omitempty"`
	PageSize int `form:"pageSize" json:"pageSize,omitempty"`
}

func (p paginationQuery) offSet() int {
	return (p.Current - 1) * p.PageSize
}

func (p paginationQuery) limit() int {
	return p.PageSize
}

type orderQuery struct {
	SortField string `form:"sortField" json:"sortField,omitempty"`
	SortOrder string `form:"sortOrder" json:"sortOrder,omitempty"`
}

func (o orderQuery) order() []ent.Order {
	var order []ent.Order
	if o.SortField != "" || o.SortOrder != "" {
		if o.SortOrder == "ascend" {
			order = append(order, ent.Asc(o.SortField))
		} else {
			order = append(order, ent.Desc(o.SortField))
		}
	} else {
		order = append(order, ent.Desc("updatedAt"))
	}
	return order
}

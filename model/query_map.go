package model

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"strconv"
)

type QueryMap struct {
	qm map[string]string
}

func NeqQueryMap() QueryMap {
	return QueryMap{qm: make(map[string]string)}
}

func ParesQuery(c *gin.Context, query []string) QueryMap {
	q := QueryMap{qm: make(map[string]string)}
	q.Parse(c, query)
	return q
}

// sortField 排序字段
// sortOrder 排序方式 [descend|ascend]
func (p QueryMap) Parse(c *gin.Context, query []string) {
	pagination := []string{"current", "pageSize"}
	order := []string{"sortField", "sortOrder"}
	query = append(query, pagination...)
	query = append(query, order...)
	for _, v := range query {
		if val, ok := c.GetQuery(v); ok {
			p.qm[v] = val
		}
	}
}

// OffSet 分页查询的偏移量
func (p QueryMap) OffSet() int {
	curr, ok1 := p.GetInt("current")
	pageSize, ok2 := p.GetInt("pageSize")
	if ok1 && ok2 {
		return (curr - 1) * pageSize
	}
	return 0
}

// Limit 分页查询的Limit
func (p QueryMap) Limit() int {
	if pageSize, ok2 := p.GetInt("pageSize"); ok2 {
		return pageSize
	}
	return 0
}

// Order 查询的排序方式
func (p QueryMap) Order() []ent.OrderFunc {
	var order []ent.OrderFunc
	SortField, ok := p.GetString("sortField")
	if ok {
		if SortOrder, ok2 := p.GetString("sortOrder"); ok2 && SortOrder == "ascend" {
			order = append(order, ent.Asc(SortField))
		} else if ok2 && SortOrder == "descend" {
			order = append(order, ent.Desc(SortField))
		} else {
			order = append(order, ent.Asc(SortField))
		}
	}
	return order
}

// GetInt 获取int类型的值
func (p QueryMap) GetInt(key string) (val int, ok bool) {
	var (
		str string
		err error
	)
	if str, ok = p.qm[key]; !ok {
		return
	}
	if val, err = strconv.Atoi(str); err != nil {
		return 0, false
	}
	return
}

// GetBool 获取bool类型的值
func (p QueryMap) GetBool(key string) (val bool, ok bool) {
	var (
		str string
		err error
	)
	if str, ok = p.qm[key]; !ok {
		return
	}
	if val, err = strconv.ParseBool(str); err != nil {
		return false, false
	}
	return
}

// GetString 获取string类型的值
func (p QueryMap) GetString(key string) (val string, ok bool) {
	val, ok = p.qm[key]
	return
}

func (p QueryMap) SetInt(key string,val int)  {
	p.qm[key]=strconv.Itoa(val)
}

func (p QueryMap) Query() map[string]string {
	return p.qm
}

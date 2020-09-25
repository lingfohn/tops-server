package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/menu"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
)

type MenuHandler struct {
	m *model.MenuManager
}

func NewMenuHandler() *MenuHandler {
	return &MenuHandler{m: model.NewMenuManager()}
}

func (p *MenuHandler) CreateMenu(c *gin.Context) (err error) {
	var m ent.Menu
	if err := c.BindJSON(&m); err != nil {
		return svcerr.Error(err, svcerr.JSONBindErr)
	}
	var menu *ent.Menu
	menu, err = p.m.CreateMenu(m)
	if err != nil {
		return
	}
	c.Set("data", menu)
	return nil
}

func (p *MenuHandler) QueryMenu(c *gin.Context) (err error) {
	query := []string{menu.FieldLevel, menu.FieldPath, "withParent", "withSubmenus"}
	qm := model.ParesQuery(c, query)
	count, menus, err := p.m.QueryAll(qm)
	if err != nil {
		return err
	}
	c.Set("data", menus)
	c.Set("total", count)
	return nil
}

func (p *MenuHandler) ModifyMenu(c *gin.Context) (err error)  {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err!=nil{
		return err
	}
	var mo ent.Menu
	if err = c.ShouldBind(&mo);err!=nil{
		return
	}
	var menu *ent.Menu
	menu,err = p.m.Update(idn,mo)
	if err!=nil{
		return
	}
	c.Set("data",menu)
	return
}
func (p *MenuHandler) DeleteMenu(c *gin.Context) (err error) {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err = p.m.Delete(idn); err != nil {
		return
	}
	c.Set("message", "删除角色成功")
	return nil
}

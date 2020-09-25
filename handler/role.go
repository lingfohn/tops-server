package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/role"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
)

type RoleHandler struct {
	rm *model.RoleManager
}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{rm: model.NewRoleManager()}
}

func (r *RoleHandler) RoleList(c *gin.Context) (err error) {
	qm := model.ParesQuery(c, role.Columns)
	count, roles, err := r.rm.Query(qm)
	if err != nil {
		return
	}
	c.Set("total", count)
	c.Set("data", roles)
	c.Set("query", qm.Query())
	return
}

func (r *RoleHandler) CreateRole(c *gin.Context) (err error) {
	var ro ent.Role
	if err := c.BindJSON(&ro); err != nil {
		return svcerr.Error(err, svcerr.JSONBindErr)
	}
	var role *ent.Role
	if role, err = r.rm.CreateRole(&ro); err != nil {
		return
	}
	c.Set("data", role)
	return
}

func (r *RoleHandler) DeleteRole(c *gin.Context) (err error) {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err = r.rm.Delete(idn); err != nil {
		return
	}
	c.Set("message", "删除角色成功")
	return nil
}

func (r *RoleHandler) ModifyRole(c *gin.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	var ro ent.Role
	if err = c.ShouldBindJSON(&ro); err != nil {
		return
	}
	var role *ent.Role
	role, err = r.rm.Update(id, ro)
	if err != nil {
		return err
	}
	c.Set("data", role)
	return nil
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/permission"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
	"strconv"
)

type PermissionHandler struct {
	pm *model.PermissionManager
}

func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{
		pm: model.NewPermissionManager()}
}

func (h *PermissionHandler) CreatePermission(c *gin.Context) (err error) {
	var p ent.Permission
	if err := c.BindJSON(&p); err != nil {
		return svcerr.Error(err, svcerr.JSONBindErr)
	}
	var perm *ent.Permission
	perm, err = h.pm.CreatePermission(p)
	if err != nil {
		return
	}
	c.Set("data", perm)
	return nil
}

func (h *PermissionHandler) QueryPermission(c *gin.Context) (err error) {
	qm := model.ParesQuery(c, permission.Columns)
	count, perms, err := h.pm.Query(qm)
	if err != nil {
		return err
	}
	c.Set("total", count)
	c.Set("data", perms)
	c.Set("query", qm.Query())
	return nil
}

func (h *PermissionHandler) ModifyPermission(c *gin.Context) (err error) {
	id := c.Param("id")
	idn, _ := strconv.Atoi(id)
	var perm ent.Permission
	if err = c.ShouldBindJSON(&perm); err != nil {
		return
	}
	permission, err := h.pm.Update(idn, perm)
	if err != nil {
		return
	}
	c.Set("data", permission)
	return nil
}

func (h *PermissionHandler) DeletePermission(c *gin.Context) (err error) {
	id := c.Param("id")
	idn, _ := strconv.Atoi(id)
	h.pm.Delete(idn)
	c.Set("message", "删除权限成功")
	return
}

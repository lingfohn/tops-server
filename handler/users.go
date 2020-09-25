package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/user"
	"github.com/lingfohn/lime/model"
	"github.com/lingfohn/lime/svcerr"
)

var (
	CreateUserErr = svcerr.NewSvcErrCode(112233, "用户创建失败")
)

type UserHandler struct {
	um *model.UserManager
}

func NewUserHandler() *UserHandler {
	return &UserHandler{um: model.NewUserManager()}
}

func (h *UserHandler) CreateUser(c *gin.Context) (err error) {
	var u *ent.User
	if err := c.Bind(&u); err != nil {
		return svcerr.Error(err, svcerr.JSONBindErr)
	}
	if u, err = h.um.Create(u); err != nil {
		return svcerr.Error(err, CreateUserErr)
	}
	c.Set("data", u)
	return
}

func (h *UserHandler) ListUser(c *gin.Context) error {
	qm := model.ParesQuery(c, user.Columns)
	//count, users, err := h.um.Query(qm)
	count, users, err := h.um.Query(qm)
	if err != nil {
		return err
	}
	c.Set("total", count)
	c.Set("data", users)
	c.Set("query", qm.Query())
	return nil
}

func (h *UserHandler) ModifyUser(c *gin.Context) error {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	var u ent.User
	if err := c.BindJSON(&u); err != nil {
		return err
	}
	user, err := h.um.UpdateUser(idn, u)
	if err != nil {
		return err
	}
	c.Set("result", user)
	return nil
}

func (h *UserHandler) DeleteUser(c *gin.Context) error {
	id := c.Param("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if err := h.um.Delete(idn); err != nil {
		return err
	}
	c.Set("message", "删除用户成功")
	return nil
}

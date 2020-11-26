package handler

import "github.com/gin-gonic/gin"

type InstanceHandler struct {

}

func NewInstanceHandler() *InstanceHandler {
	return &InstanceHandler{}
}

func (i *InstanceHandler)CreateInstance(c *gin.Context)(err error)  {
	return nil
}

func (i *InstanceHandler)QueryInstance(c *gin.Context)(err error)  {
	return nil
}
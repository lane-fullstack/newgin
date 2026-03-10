package ctx

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

func New(c *gin.Context) *Context {
	return &Context{Context: c}
}
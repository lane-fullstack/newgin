package ctx

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (c *Context) Success(data interface{}, total ...int64) {
	if len(total) > 0 {
		resp := gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"list":  data,
				"total": total[0],
			},
		}
		c.JSON(200, resp)
		return
	}

	resp := gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	}
	c.JSON(200, resp)

}

func (c *Context) Error(msg string) {

	c.JSON(200, Response{
		Code: 1,
		Msg:  msg,
	})
}

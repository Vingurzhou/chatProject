package ctrl

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Addfriend(c *gin.Context) {
	err := errors.New("")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "好友添加成功",
			"data": nil,
		})
	}
}

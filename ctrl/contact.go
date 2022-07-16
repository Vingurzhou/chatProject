package ctrl

import (
	"chatProject/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var contactService service.ContactService

func Addfriend(c *gin.Context) {

	userid := c.PostForm("userid")
	distid := c.PostForm("distid")

	nuserid, _ := strconv.Atoi(userid)
	ndistid, _ := strconv.Atoi(distid)
	err := contactService.AddFriend(int64(nuserid), int64(ndistid))

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

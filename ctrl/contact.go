package ctrl

import (
	"chatProject/model"
	"chatProject/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var contactService service.ContactService

func AddFriend(c *gin.Context) {

	userid := c.PostForm("userid")
	dstid := c.PostForm("dstid")
	nuserid, _ := strconv.Atoi(userid)
	ndstid, _ := strconv.Atoi(dstid)
	err := contactService.AddFriend(int64(nuserid), int64(ndstid))

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

func LoadFriend(c *gin.Context) {

	userid := c.PostForm("userid")

	nuserid, _ := strconv.Atoi(userid)
	users := contactService.SearchFriend(int64(nuserid))
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"rows":  users,
		"total": len(users),
	})
}

func CreateCommunity(c *gin.Context) {

	arg := model.Community{}

	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": com,
		})
	}

}

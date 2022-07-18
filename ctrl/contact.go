package ctrl

import (
	"chatProject/model"
	"chatProject/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var contactService service.ContactService

type PageArg struct {
	//从哪页开始
	Pagefrom int `json:"pagefrom" form:"pagefrom"`
	//每页大小
	Pagesize int `json:"pagesize" form:"pagesize"`
	//关键词
	Kword string `json:"kword" form:"kword"`
	//asc：“id”  id asc
	Asc  string `json:"asc" form:"asc"`
	Desc string `json:"desc" form:"desc"`
	//
	Name string `json:"name" form:"name"`
	//
	Userid int64 `json:"userid" form:"userid"`
	//dstid
	Dstid int64 `json:"dstid" form:"dstid"`
	//时间点1
	Datefrom time.Time `json:"datafrom" form:"datafrom"`
	//时间点2
	Dateto time.Time `json:"dateto" form:"dateto"`
	//
	Total int64 `json:"total" form:"total"`
}

func (p *PageArg) GetPageSize() int {
	if p.Pagesize == 0 {
		return 100
	} else {
		return p.Pagesize
	}

}
func (p *PageArg) GetPageFrom() int {
	if p.Pagefrom < 0 {
		return 0
	} else {
		return p.Pagefrom
	}
}

func (p *PageArg) GetOrderBy() string {
	if len(p.Asc) > 0 {
		return fmt.Sprintf(" %s asc", p.Asc)
	} else if len(p.Desc) > 0 {
		return fmt.Sprintf(" %s desc", p.Desc)
	} else {
		return ""
	}
}

type ContactArg struct {
	PageArg
	Userid int64 `json:"userid" form:"userid"`
	Dstid  int64 `json:"dstid" form:"dstid"`
}

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
	var arg ContactArg
	c.Bind(&arg)
	users := contactService.SearchFriend(arg.Userid)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"rows":  users,
		"total": len(users),
	})
}

func CreateCommunity(c *gin.Context) {

	arg := model.Community{}
	c.Bind(&arg)
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

func LoadCommunity(c *gin.Context) {
	var arg ContactArg
	//如果这个用的上,那么可以直接
	c.Bind(&arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"rows":  comunitys,
		"total": len(comunitys),
	})
}
func JoinCommunity(c *gin.Context) {
	var arg ContactArg
	c.Bind(&arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	AddGroupId(arg.Userid, arg.Dstid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "",
			"data": nil,
		})
	}
}
func AddGroupId(userId, gid int64) {
	//取得node
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	//clientMap[userId] = node
	rwlocker.Unlock()
	//添加gid到set
}

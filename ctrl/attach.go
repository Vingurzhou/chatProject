package ctrl

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

const (
	AccessKeyId     = "5p2RZKnrUanMuQw9"
	AccessKeySecret = "bsNmjU8Au08axedV40TRPCS5XIFAkK"
	EndPoint        = "oss-cn-shenzhen.aliyuncs.com"
	Bucket          = "winliondev"
)

func Upload(c *gin.Context) {
	srcfile, head, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}

	//todo 获得文件后缀.png/.mp3

	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := c.PostForm("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}

	//todo 初始化ossclient
	client, err := oss.New(EndPoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
	//todo 获得bucket
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
	//todo 设置文件名称
	//time.Now().Unix()
	filename := fmt.Sprintf("mnt/%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	//todo 通过bucket上传
	err = bucket.PutObject(filename, srcfile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
	//todo 获得url地址
	url := "http://" + Bucket + "." + EndPoint + "/" + filename

	//todo 响应到前端
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": url,
	})
}

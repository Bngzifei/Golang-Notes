package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
gin 框架中采用的路由库是基于httprouter做的

API参数:可以通过Context的Param方法来获取API参数

*/

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		// 截取
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})

	// 监听端口默认为8000
	r.Run(":8000")

}

package routers

import (
	"github.com/gin-gonic/gin"
	v1 "service/routers/api/v1"
)

func Init() {

	r := gin.Default()

	v := r.Group("/api/v1")
	{
		// 登录
		v.POST("/login", v1.Login)
		// 获取 msg
		v.POST("/getmsg", v1.GetMessage)
	}

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}

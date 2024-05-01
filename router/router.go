package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/attach-file/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// attach file 表操作
		Router.POST("upload", api.Upload)
		Router.GET("download", api.Download)
		Router.POST("list", api.ListAttachFile)
	}
}

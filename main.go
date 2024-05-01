package attachfile

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/attach-file/global"
	"github.com/provider-go/attach-file/router"
	"gorm.io/gorm"
)

type Plugin struct{}

func CreatePlugin() *Plugin {
	return &Plugin{}
}

func CreatePluginAndDB(db *gorm.DB) *Plugin {
	global.DB = db
	return &Plugin{}
}

func (*Plugin) Register(group *gin.RouterGroup) {
	router.GroupApp.InitRouter(group)
}

func (*Plugin) RouterPath() string {
	return "file"
}

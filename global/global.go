package global

import (
	"github.com/provider-go/pkg/file"
	"github.com/provider-go/pkg/file/typefile"
	"github.com/provider-go/pkg/logger"
	"github.com/provider-go/pkg/smcc"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	AttachFile file.StorageFile
	SMCC       smcc.SMCC
)

func InstanceStorageFile() {
	// 地址改成从配置中心读取
	addr, err := SMCC.GetConfig("ipfs.addr")
	if err != nil {
		logger.Error("init:", "step", "GetConfig", "err", err)
	}
	cfg := typefile.ConfigStorageFile{Endpoints: addr}
	AttachFile = file.NewStorageFile("ipfs", cfg)
}

package global

import (
	"github.com/provider-go/pkg/file"
	"github.com/provider-go/pkg/file/typefile"
	"gorm.io/gorm"
)

var DB *gorm.DB

var AttachFile file.StorageFile

func init() {
	cfg := typefile.ConfigStorageFile{Endpoints: "192.168.0.103:5001"}
	AttachFile = file.NewStorageFile("ipfs", cfg)
}

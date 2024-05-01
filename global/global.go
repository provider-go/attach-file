package global

import (
	shell "github.com/ipfs/go-ipfs-api"
	"gorm.io/gorm"
)

var DB *gorm.DB

var AttachFile *shell.Shell

func init() {
	AttachFile = shell.NewShell("192.168.0.103:5001")
}

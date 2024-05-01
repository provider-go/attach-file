package models

import (
	"github.com/provider-go/attach-file/global"
	"time"
)

type AttachFile struct {
	Hash       string    `json:"hash" gorm:"column:hash;type:varchar(40);primary_key;comment:'主键'"`
	FilePath   string    `json:"filePath" gorm:"column:file_path;type:varchar(200);not null;default:'';comment:文件保存路径"`
	FileName   string    `json:"fileName" gorm:"column:file_name;type:varchar(200);not null;default:'';comment:原始文件名称"`
	FileSize   int64     `json:"fileSize" gorm:"column:file_size;not null;default:0;comment:文件大小"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `json:"update_time" gorm:"autoCreateTime;comment:更新时间"`
}

func CreateAttachFile(hash, filePath, fileName string, fileSize int64) error {
	return global.DB.Table("attach_files").Select("hash", "filePath", "fileType", "fileSize").
		Create(&AttachFile{Hash: hash, FilePath: filePath, FileName: fileName, FileSize: fileSize}).Error
}

func ListAttachFile(pageSize, pageNum int) ([]*AttachFile, int64, error) {
	var rows []*AttachFile
	//计算列表数量
	var count int64
	global.DB.Table("attach_files").Count(&count)

	if err := global.DB.Table("attach_files").Order("CreateTime desc").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func ViewAttachFile(hash string) (*AttachFile, error) {
	row := new(AttachFile)
	if err := global.DB.Table("attach_files").Where("hash = ?", hash).First(&row).Error; err != nil {
		return nil, err
	}
	return row, nil
}

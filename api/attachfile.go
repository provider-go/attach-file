package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/attach-file/global"
	"github.com/provider-go/attach-file/models"
	"github.com/provider-go/pkg/ioput"
	"github.com/provider-go/pkg/ipfs"
)

func Upload(ctx *gin.Context) {
	// 接收单文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "接收单文件错误~")
		return
	}
	// 文件上传至ipfs
	src, err := file.Open()
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "接收文件打开错误~")
		return
	}

	hash, err := ipfs.UploadIPFS(global.AttachFile, src)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "文件上传至ipfs错误~")
		return
	}
	// 上传记录入库
	err = models.CreateAttachFile(hash, "", file.Filename, file.Size)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "上传文件入库错误~")
	} else {
		ioput.ReturnSuccessResponse(ctx, hash)
	}
}

func Download(ctx *gin.Context) {
	hash := ctx.Query("hash")
	s, err := ipfs.CatIPFS(global.AttachFile, hash)
	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "从ipfs下载错误~")
		return
	}
	ctx.Writer.Write(s)
}

func ListAttachFile(ctx *gin.Context) {
	j := make(map[string]interface{})
	_ = ctx.BindJSON(&j)
	pageSize := ioput.ParamToInt(j["pageSize"])
	pageNum := ioput.ParamToInt(j["pageNum"])
	list, total, err := models.ListAttachFile(pageSize, pageNum)

	if err != nil {
		ioput.ReturnErrorResponse(ctx, 9999, "系统错误~")
	} else {
		res := make(map[string]interface{})
		res["records"] = list
		res["total"] = total
		ioput.ReturnSuccessResponse(ctx, res)
	}
}

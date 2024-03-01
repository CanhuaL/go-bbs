package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_bbs/logic"
	"go_bbs/models"
	"net/http"
	"strconv"
	"time"
)

func CreateCommentHandler(c *gin.Context) {
	p := new(models.ParamComment)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 从请求中获取上传的文件
	file, err := c.FormFile("picture")

	// 生成图片URL
	imageName := file.Filename
	imageURL := generateFileName(imageName)

	// 保存图片到服务器
	imagePath := fmt.Sprintf("./temp/%s", imageName)
	if err = c.SaveUploadedFile(file, imagePath); err != nil {
		ResponseError(c, http.StatusInternalServerError)
		return
	}

	//  新建评论,将评论照片url插入数据库，并上传到oss
	now := time.Now()
	p.CreateTime = now
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.PostId = id
	if err := logic.CreateComment(p, imageURL, imageName, imagePath); err != nil {
		ResponseError(c, CodeCreateCommentErr)
		fmt.Println(err)
		return
	}

	ResponseSuccess(c, nil)
}

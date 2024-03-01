package logic

import (
	"go.uber.org/zap"
	"go_bbs/dao/mysql"
	"go_bbs/dao/oss"
	"go_bbs/models"
	"os"
)

func CreateComment(p *models.ParamComment, imageURL, imageName, imagePath string) (err error) {
	// 上传图片到阿里云 OSS
	if err = oss.UploadAvatarToOSS(imagePath, imageName); err != nil {
		zap.L().Error("Failed to upload image to OSS", zap.Error(err))
		return
	}

	// 插入评论到到mysql中
	if err = mysql.CreateComment(p, imageURL); err != nil {
		return err
	}

	// 删除服务器上的临时图片文件
	if err = os.Remove(imagePath); err != nil {
		zap.L().Error("Failed to delete image file", zap.Error(err))
	}
	return nil
}

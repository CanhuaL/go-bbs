package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_bbs/logic"
	"go_bbs/models"
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
	//  新建评论
	now := time.Now()
	p.CreateTime = now
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.PostId = id
	if err := logic.CreateComment(p); err != nil {
		ResponseError(c, CodeCreateCommentErr)
		fmt.Println(err)
		return
	}

	ResponseSuccess(c, nil)
}

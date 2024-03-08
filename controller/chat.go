package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_bbs/dao/mysql"
	"go_bbs/logic"
	"net/http"
	"strconv"
)

func PrivateChatHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	//p := new(models.ParamPrivateChat)
	//if err := c.ShouldBindJSON(p); err != nil {
	//	// 请求参数有误，直接返回响应
	//	zap.L().Error("PrivateChat with invalid param", zap.Error(err))
	//	// 判断err是不是validator.ValidationErrors 类型
	//	errs, ok := err.(validator.ValidationErrors)
	//	if !ok {
	//		ResponseError(c, CodeInvalidParam)
	//		return
	//	}
	//	ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
	//	return
	//}
	//senderIdStr := c.Param("senderId")
	//senderId, err := strconv.ParseInt(senderIdStr, 10, 64) // 转换为int64
	//if err != nil {
	//	return
	//}
	// 使用c.Param获取路由参数
	senderIDStr := c.Query("sender_id")
	receiverIDStr := c.Query("receiver_id")
	// 将参数字符串转换为int64
	senderID, err := strconv.ParseInt(senderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sender_id"})
		return
	}

	receiverID, err := strconv.ParseInt(receiverIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid receiver_id"})
		return
	}
	// 2. 业务处理
	if err := logic.PrivateChat(senderID, receiverID, c.Writer, c.Request); err != nil {
		zap.L().Error("logic.DeleteFriend failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

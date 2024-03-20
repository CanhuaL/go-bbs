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

// CheckPrivateChatHandler 查询用户间私聊信息
func CheckPrivateChatHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	senderIdStr := c.Param("sender_id")
	senderId, err := strconv.ParseInt(senderIdStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	receiverIdStr := c.Param("receiver_id")
	receiverId, err := strconv.ParseInt(receiverIdStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.CheckPrivateChat(senderId, receiverId, page, size)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

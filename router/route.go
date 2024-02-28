package router

import (
	"go_bbs/controller"
	"go_bbs/logger"
	"go_bbs/middlewares"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "go_bbs/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.DebugMode {
		gin.SetMode(gin.DebugMode) // gin设置成Debug模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录默认通过用户名和密码
	v1.POST("/login", controller.LoginHandler)
	//  手机号码登录
	v1.POST("/phone_login", controller.PhoneLoginHandler)
	//  邮箱登录
	v1.POST("/email_login", controller.EmailLoginHandler)
	//  手机验证码登录
	v1.POST("/sms_login", controller.SMSLoginHandler)
	//  发送验证码
	v1.POST("/send_code", controller.SendCode)
	// 根据时间或分数获取帖子列表
	v1.GET("/posts2", controller.GetPostListHandler2)
	//  获取帖子列表的处理函数
	v1.GET("/posts", controller.GetPostListHandler)
	//  查询到所有的社区,以列表的形式返回
	v1.GET("/community", controller.CommunityHandler)
	//  根据ID查询社区详情
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	//  查询帖子详情
	v1.GET("/post/:id", controller.GetPostDetailHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		//  创建帖子
		v1.POST("/post", controller.CreatePostHandler)
		//  帖子评论相关

		//  新增帖子评论
		v1.POST("/create_comment/:id", controller.CreateCommentHandler)
		//  删除帖子评论
		v1.DELETE("/delete_comment")
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

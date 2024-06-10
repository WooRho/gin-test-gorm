package main

import (
	"gin-test-gorm/config"
	"gin-test-gorm/model"
	"gin-test-gorm/server"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		err error
	)
	// init
	config.Init()

	// 1.创建路由
	r := gin.Default()

	//err = r.SetTrustedProxies([]string{strconv.Itoa(config.Config..)})
	//if err != nil {
	//	return
	//}

	initServer(r)

	err = r.Run(":" + config.Config.Web.Port)
	if err != nil {
		return
	}
}

func initServer(r *gin.Engine) {
	// 初始化数据库
	model.Migrate(model.GetDb())

	baseGroup := r.Group("")

	adminAuth := baseGroup.Group("", server.AuthAdmin)
	//// 登录登出
	//{
	baseGroup.POST("login", server.Login)
	//	v1Admin.GET("初始化路由和超管权限", "syncRouterToDB", system.SyncRouterToDB)
	//	// 手机验证码登录
	//	v1Admin.POST("发送验证码注册", "sendVerificationCodeEmail", system.SendVerificationCodeEmail)
	//	v1Admin.POST("手机验证码注册", "phoneVerificationCodeLogin", system.PhoneVerificationCodeLogin)
	//
	//	v1Admin.POST("登出", "logout", system.Logout)
	//
	//	adminAuth.GET("获取用户信息", "information", system.GetLoginInformation)
	//	adminAuth.GET("获取菜单", "getMenus", system.GetMenus)
	//	// 查询登录日志
	//	adminAuth.GET("获取登录日志", "loginLog", system.GetLoginLogList)
	//	adminAuth.GET("获取登录地", "loginSite", system.GetLoginSite)
	//}
	adminAuth.POST("create", server.Create)
}

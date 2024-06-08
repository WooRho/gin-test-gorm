package main

import (
	"gin-test-gorm/config"
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

}

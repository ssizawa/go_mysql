package main

import (
	//Ginの実装
	"github.com/gin-gonic/gin"

	"github.com/ssizawa/go_mysql/server/src/controller"
)

func main() {

	router := gin.Default()

	//html/css/jsファイルのディレクトリをロード
	router.LoadHTMLGlob("../../frontend/templates/*")
	router.Static("../../frontend/assets", "./assets")

	//routerを渡す
	controller.Router(router)

	//ポートを指定して実行
	router.Run(":80")
}

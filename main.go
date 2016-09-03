package main

import (
	"./handler"
	"github.com/gin-gonic/gin"
)

func main() {
  // 面倒な初期処理を一括で済ませてくれる
	router := gin.Default()

  // ルーティング
  // どのURLにアクセスした時どういった処理を行うか決めること
  // routerはポインタなので変更が伝わる
	handler.UserRoutes(router)

  // サーバー起動
	router.Run(":8080")
}

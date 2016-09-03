package handler

import (
	"../db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// どんな方法でどのURLにアクセスされた時、何をするかを定義していく
func UserRoutes(router *gin.Engine) {
	// GETは普通にブラウザからアクセスすればいい
	// ブラウザから http://localhost:8080/users でアクセスされた時
	router.GET("/users", getAllUsers)

	// ブラウザから http://localhost:8080/users?id=*********** でアクセスされた時
	router.GET("/users:id", getUser)

	// POSTで http://localhost:8080/users にアクセスされた時
	router.POST("/users", createUser)

	// PATCHで http://localhost:8080/users にアクセスされた時
	router.PATCH("/users", updateUser)

	// DELETEで http://localhost:8080/users にアクセスされた時
	router.DELETE("/users", deleteUser)
}

// 全てのユーザのBSONをデータベースから取得し、JSONに変換してアクセス元に返す
func getAllUsers(context *gin.Context) {
	users, error := db.User{}.FindAll()
	if error != nil {
		// ステータスコードだけ返す
		// JSONデータはなし
		// status code: 404 Not found
		context.Status(404)
		return
	}

	// 取得したデータをJSONに変換してレスポンスとして返す
	// status code: 200 OK
	context.JSON(200, users)
}

// URLのGETパラメータ(URLの最後の?以降)で指定されたidを持つユーザのBSONをデータベースから取得し、JSONに変換してアクセス元に返す
func getUser(context *gin.Context) {
	// URLに付与されたパラメータを文字列として取得
	stringID := context.Param("id")

	// 取得した文字列がObjectId型に変換できるか確認
	if !bson.IsObjectIdHex(stringID) {
		// status code: 400 Bad request
		context.Status(400)
		return
	}
	// bson.ObjectId型に変換
	id := bson.ObjectIdHex(stringID)

	user, error := db.User{}.Find(id)
	if error != nil {
		// status code: 404 Not found
		context.Status(404)
		return
	}

	// status code: 200 OK
	context.JSON(200, user)
}

// POSTでリクエストされたJSONデータをBSONに変換してデータベースに保存
func createUser(context *gin.Context) {
	user := new(db.User)
	if error := context.BindJSON(user); error != nil {
		// リクエストデータのJSONをUser型の変数に変換できないということは
		// リクエストデータがおかしい
		// status code: 400 Bad request
		log.Println("failed to bind JSON")
		context.Status(400)
		return
	}
	// IDを新しく生成
	user.ID = bson.NewObjectId()

	if error := user.Save(); error != nil {
		// status code: 400 Bad request
		log.Println("failed to save user")
		context.Status(400)
		return
	}

	// status code: 201 Created
	context.Status(201)
}

// PATCHでリクエストされたデータのidに一致するデータをデータベースから探し、それをリクエストされたデータの内容で書き換える
func updateUser(context *gin.Context) {
	user := new(db.User)
	if error := context.BindJSON(user); error != nil {
		// status code: 400 Bad request
		context.Status(400)
		return
	}

	if error := user.Update(); error != nil {
		// status code: 400 Bad request
		context.Status(400)
		return
	}

	// status code: 200 OK
	context.JSON(200, user)
}

// DELETEでリクエストされたデータのidに一致するデータをデータベースから削除する
func deleteUser(context *gin.Context) {
	user := new(db.User)
	if error := context.BindJSON(user); error != nil {
		// status code: 400 Bad request
		context.Status(400)
		return
	}

	if error := user.Delete(); error != nil {
		// status code: 404 Not found
		context.Status(404)
		return
	}

	// 削除の成功のみステータスコードで伝える
	// status code: 204 No Content
	context.Status(204)
}

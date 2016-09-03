package db

import "gopkg.in/mgo.v2/bson"

// 保存する関数
func (user *User) Save() error {
  // バリデーション
  if error := user.isValid(); error != nil {
    return error
  }

  // usersコレクションに接続
  session, collection := connect("users")
  // この関数の処理が終わるか中断した時に必ず実行する
  // データベースとの接続を切断する
  defer session.Close()

  // userのデータをusersコレクションに保存
  error := collection.Insert(user)
  return error
}

func (user *User) Update() error {
  if error := user.isValid(); error != nil {
    return error
  }

  session, collection := connect("users")
  defer session.Close()

  // 第一引数に指定したIDのドキュメントの内容をi
  // 第二引数に指定したUser型の変数の内容で書き換える
  error := collection.UpdateId(user.ID, user)
  return error
}

func (user *User) Delete() error {
  session, collection := connect("users")
  defer session.Close()

  // 指定したIDのドキュメントを削除
  error := collection.RemoveId(user.ID)
  return error
}

func (_ User) Find(id bson.ObjectId) (*User, error) {
  session, collection := connect("users")
  defer session.Close()

  // 検索結果を代入するためのポインタ変数
  user := new(User)
  // 検索結果を取得
  query := collection.FindId(id)
  // 検索結果から1件取得(idはユニークなので1件しかない)
  error := query.One(user)

  // error := collection.FindId(id).One(user) としても良い

  return user, error
}

func (_ User) FindAll() (*[]User, error) {
  session, collection := connect("users")
  defer session.Close()

  users := new([]User)
  // 条件なしで取得(つまり全件取得)
  query := collection.Find(nil)
  // 検索結果を全て取得(スライス)
  error := query.All(users)
  return users, error
}

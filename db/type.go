package db

import (
  // 独自のエラーを定義できる
  // といっても表示されるエラー文を決められるだけ
  "errors"
  "gopkg.in/mgo.v2/bson"
)

// bsonはデータベースに保存する時の名前、jsonはJSONに変換した時の名前
type User struct {
  // IDはデフォルトではMongoDBが自動で作成してドキュメントに含める
  // それをGoで制御するためには、変数の型をbson.ObjectIdとし、bsonを_idにする
  ID   bson.ObjectId `bson:"_id" json:"id"`
  Name string        `bson:"name" json:"name"`
  Age  int           `bson:"age"  json:"age"`
}

// バリデーション(データが想定したものでなければエラー)
func (user *User) isValid() error {
  // 名前は空文字ではない
  if user.Name == "" {
    return errors.New("InvalidMemberError at User.Name")
  }
  // 年齢は0以上
  if user.Age < 0 {
    return errors.New("InvalidMemberError at User.Age")
  }
  return nil
}

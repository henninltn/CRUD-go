package db

// / がある場合 / 以下の名前を使う
// mgo.v2は長いのでmgoだけで使えるようにする
import mgo "gopkg.in/mgo.v2"

func connect(collectionName string) (*mgo.Session, *mgo.Collection) {
  // MongoDBのサーバーに接続(データベースはサーバーのコンピュータで動かすソフト)
  // ２つ目の返り値はerror型だが、使わないので_と書いておく
  session, _ := mgo.Dial("localhost")

  // "sample"という名前のデータベースに接続
  db := session.DB("sample")
  // collectionNameという名前のコレクションに接続
  collection := db.C(collectionName)

  // collection := session.DB("sample").C(collectionName) としても良い

  return session, collection
}

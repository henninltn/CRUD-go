CRUD-back
=========

簡単なユーザ管理のサーバー

## Description
- できること
  - 全てのユーザのデータの取得
  - 特定ユーザのデータの取得
  - ユーザの作成
  - ユーザの情報の書き換え
  - ユーザの削除

## Installation & Useag
```
$ sudo pacman -Syu
> password: your password
$ sudo pacman -S mongodb

$ go get github.com/gin-gonic/gin

$ mkdir -p $GOPATH/src/gitlab.com/hennin
$ cd $GOPATH/src/gitlab.com/hennin
$ git clone https://gitlab.com/hennin/CRUD-back.git
$ cd CRUD-back

$ systemctl start mongodb
$ go run main.go
```

#### GET
ブラウザからhttp://localhost:8080/users にアクセス

```null```とだけ表示されるはず


もう一個ターミナル起動しとく


#### POST
```
$ cd $GOPATH/src/gitlab.com/hennin/CRUD-back/test_sh
$ ./post_test.sh http://localhost:8080/users 1 '{"name": "ichiro", "age": 20}'
$ ./post_test.sh http://localshot:8080/users 1 '{"name": "jiro", "age": 19}'
```

ブラウザで http://localhost:8080/users にアクセス

```
[{"id":"何か文字列1","name":"ichiro","age":20"},{"id":"何か文字列2","name":"jiro","age":19}]
```

と表示されるはず


#### もっかいGET
ブラウザで http://localhost:8080?id=何か文字列1 にアクセス

```
{"id":"何か文字列","name":"ichiro","age":20"}
```

と表示されるはず


#### PATCH
```
$ ./patch_test.sh http://localhost:8080/users 1 '{"id":"何か文字列","name": "ichiro", "age": 21}'
```

```
[{"id":"何か文字列1","name":"ichiro","age":21"},{"id":"何か文字列2","name":"jiro","age":19}]
```


#### DELETE
```
$ ./delete_test.sh http://localhost:8080/users 1 '{"id":"何か文字列","name": "ichiro", "age": 21}'
```

```
[{"id":"何か文字列2","name":"jiro","age":19}]
```

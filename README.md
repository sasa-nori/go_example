# Go練習用

## 環境構築

|種別|対象の環境|
|:--:|:--:|
|OS| macOS Mojave |
|Go| 1.3|

### 自分でやる編

1. [ここから](https://golang.org/dl/)最新バージョンをダウンロード
1. インストーラを使ってインストール
1. 作業用に$HOME以下にディレクトリを作成し、その直下にgoディレクトリを作成
1. 下記PATHを通す(`.bash_profile`)
    * `export GOPATH=$HOME/{作った作業用フォルダ}/go`
    * `export PATH=$PATH:$GOPATH/bin`
    * `export PATH=$PATH:/usr/local/go/bin`
1. `$GOPATH` 以下に `src` ディレクトリ作成
1. `src` へ移動してクローン
1. クローンしたディレクトリに移動
1. `go run main.go` で `main() {}` が実行されます
1. `Hello, World!` と表示されるはずです

### 自動でやってもらう編

1. これを実行 `sh -c "$(curl -fsSL https://raw.githubusercontent.com/noriyuki-sasagawa/go_example/master/setup.sh)"`
    * .shファイルの中身は[こちら](./setup.sh)
1. `Hello, World!` と表示されるはずです

## ブログ記事で取り扱ったファイル一覧

### [【第1回】Go言語（Golang）入門～環境構築編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-environment-1)

* [hello.go](./hello/hello.go)

### [【第2回】Go言語（Golang）入門～Firestore導入編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-firestore)

* [firestore.go](./firestore/firestore.go)

### [【第3回】Go言語（Golang）入門～Firestoreデータ操作編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-firestore-data-manipulator)

[firestore.go](./firestore/firestore.go)

### 4回目

[echo.go](./echo/echo.go)

### 5回目

[api.go](./api/api.go)

### 6回目

[twitter.go](./twitter/twitter.go)

### 7回目

[search.go](./search/search.go)

### 8回目

[setup.sh](./setup.sh)

### 9回目

[objectbox.go](./ob/objectbox.go)

# 練習用

## 環境構築(Mac)

1. [ここから](https://golang.org/dl/)最新バージョンをダウンロード
1. インストーラを使ってインストール
1. 作業用に$HOME以下にディレクトリを作成し、その直下にgoディレクトリを作成
1. 下記PATHを通す(`.bash_profile`)
    * `export GOPATH=$HOME/{作った作業用フォルダ}/go`
    * `export PATH=$PATH:$GOPATH/bin`
    * `export PATH=$PATH:/usr/local/go/bin`
1. `$GOPATH` 以下に `src` ディレクトリ作成
1. [dep](https://golang.github.io/dep/docs/installation.html#macos)をインストール
    * `$ brew install dep | brew upgrade dep`
1. `$GOPATH/src`に移動してこちらをクローン
1. `$ dep ensure` を実行
1. `$ go run` で `main()` が実行されます
1. `http://localhost:1323/hello` で `Hello, World!` と表示されるはずです

## 1日目

[hello.go](./hello/hello.go)

## 2日目

[firestore.go](./firestore/firestore.go)

## 3日目

[firestore.go](./firestore/firestore.go)

## 4日目

[echo.go](./echo/echo.go)

## 5日目

[api.go](./api/api.go)

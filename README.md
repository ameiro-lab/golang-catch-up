### 1. モジュールを作成する
モジュール名は 一意であることが望ましい ため、通常は **ドメイン名 + プロジェクト名** や、GitHub などにコードを公開する場合は **リポジトリの URL** に一致させるのが一般的です。
```sh
# Go モジュールを初期化するためのコマンド
go mod init モジュール名
```
一般的な命名規則： _ より - の方が好まれることが多い。例：golang-practice


### 2. モジュール内にパッケージを作成する
Go では、モジュールの中に複数のパッケージ を作成できます。  
まず、開発する機能に応じて ディレクトリ構成を考えましょう。
```
root/   ← モジュールのルートディレクトリ
  ├── go.mod   ← モジュール情報
  ├── main.go  ← メインプログラム
```
小規模なプロジェクトなら、上記のように main.go を作るだけでもOK！もしコードを整理したい場合は、下記のようにパッケージを分割してディレクトリを作成すること。
```
root/
  ├── go.mod
  ├── main.go
  ├── greeting/         ← カスタムパッケージ
  │   ├── hello.go
  │   ├── goodmorning.go
```

### 3. ソースを記述する
```go
// greeting/hello.go

package greeting

import "fmt"

func Hello() {
    fmt.Println("Hello, World!")
}
```
```go
// main.go

package main

import (
    "myapp/greeting"  // パッケージをインポート
)

func main() {
    greeting.Hello()
}
```
実行：  `go run main.go`
```sh
# main.go を実行コマンド
go run main.go
```

### 4. 必要なパッケージを追加する
Go の標準ライブラリ以外のパッケージを使う場合は go get でインストールします。
```sh
# Go パッケージをインストールするコマンド
go get github.com/fatih/color
```
この操作により go.mod と go.sum に依存関係が追加されます。


### 単体テストをする
Go では 標準のテストフレームワーク を使って、すぐにテストを実行できます。
テスト用の `*_test.go` ファイルを作成し、`go test` コマンドで実行できます。
```go
// greeting/hello_test.go

package greeting

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, World!"
    result := "Hello, World!" // 実際の処理に置き換える
    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}
```

### おまけ. Goの標準ライブラリの紹介




### おまけ②：gRPC
gRPCはオープンソースの **RPCフレームワーク **。マイクロサービスを作るならgRPCが有用。異なる言語間（Go⇔Javaなど）で通信しやすくできる。または、Go ⇔ Go のマイクロサービスでも使うと便利。  

**マイクロサービスとは？**  
マイクロサービスは「小さな独立したサービスの集合体」。なので、各サービスは独立して動作しており、他のサービスのデータベースには直接アクセスできない。  
データが必要ならば、API（REST / GraphQL）やRPC（gRPC / Thrift） を使って他のサービスにリクエストを送る必要がある。  
※Goのみで完結するアプリ（例えば、バックエンドが1つのサーバーで動作するモノリシックな構成）ならば、gRPCやRPCは不要。
代わりに、普通の関数呼び出しやGoのHTTPサーバー(net/http)を使えば十分。  

**結論**  
Spring Bootでログイン認証を行い、ログイン後の機能はGoのバックエンドで処理する という構成することもできる！

### 例外処理について
Goは 例外処理（Exception Handling） を使用しない言語。代わりに、エラーチェック（Error Handling） をコードの中で明示的に行うスタイルが取られているので、if ~ else ではなくて、関数の戻り値としてエラー型を返すのが一般的。
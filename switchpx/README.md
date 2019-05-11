# features
ターミナルを起動させたときに、 .gitconfigや, .curlrcに設定されているProxyを切り替えてくれる優れものです。  
プロキシサーバ下の大学や職場と、それが無い自宅の両方でPCを使うとき、.gitconfigなどのProxyの行をいちいちコメントアウトなどをしていませんか?  
それはGo言語さんのバイナリに勝手に行ってもらいましょう。  
<br><br><br>

## Example
プロキシは普通以下のように書かれていると思います。  
```
proxy=<プロキシ>:<ポート>
```
<br>

その行を、ターミナルが起動したときに自動でコメントアウトしてくれます
```
# proxy=<プロキシ>:<ポート>
```
<br>

もしくは、プロキシ下のネットワークにいないときは、コメントアウトを取り除いてくれます。  
<br><br><br>

## Dependency
go 1.11.1  
zsh  
<br><br>

## Usage
まず、プロキシ環境下のネットワークアドレスを求める必要があります。  
手計算やどこかのサイトで求められる方はそちらを。それが厳しかったら、fproxyのGetNetAddr()さんを使いましょう。   
以下のようなコードでネットワークアドレスを求められます。  
```go
package main

import (
    "fmt"
    "switchpx/fproxy"
)

func main() {
    var actual fproxy.Actual
    s, _ := actual.GetNetAddr()
	fmt.Println(s.String())
}
```

<br><br>
次に、設定ファイルを生成します。  
一度、switchpxのバイナリを実行します。  
```
./switchpx
```
<br>

そうすると、  
```
自動コメントアウトに失敗しました。 read test: is a directory
```
というコメントが出ますが、設定ファイルの作成には成功しました。  
~/.sfp/config.jsonという設定ファイルができていると思います。  
<br><br>

このconfig.jsonファイルに、.gitconfigやzshrc,curlrcなど、自動で切り替えてほしいファイルのパスを書き込みます。  
例:"filepath": "/User/ssab/.gitconfig"  
(現在~/.gitconfigなどの書き方に対応していません)  
<br>

次に、

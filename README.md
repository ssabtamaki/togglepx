# features
シェルを起動させたときに、 .gitconfigや, .curlrcに設定されているProxyを切り替えてくれる優れものです。  
大学と自宅の両方でPCを使うとき、.gitconfigなどのProxyの行をいちいちコメントアウトなどをしていませんか?  
それはGo言語さんに勝手に行ってもらいましょう。  

#Dependency
go 1.11.1
zsh

# Usage
まず、プロキシ環境下のネットワークアドレスを求める必要があります。  
手計算やどこかのサイトで求められる方はそちらを、それが厳しかったら、このfproxyパッケージをつかって求めましょう.  
以下のようなコードでネットワークアドレスを求められます。  
第1引数にネットワークアドレス、第2引数にerr値が渡されます。  
```go
package main

import (
	"fmt"
	"stepupgo/fproxy"
)

func main() {
	fmt.Println(fproxy.GetNetIPv4().string())
}
```

そしたら、main.goのconstのunivIPの値を、各自のネットワークアドレスに書き換えます。  
次に、対象のファイルを指定しましょう。constの値を各自のファイルに割り当てます。  
一応失敗してファイルが消えたら怖いので、一度copyをかけておいてください。

最後に、go buildしてバイナリを生成し、zshにバイナリのPATHを記述してあげます。  
 ```
~/go/src/stepupgo/stepupgo
``` 
こんなかんじで記述します。各々のPATHを指定してあげてください。

以上で、出来ると思います。ただ、zshに記述したあとはきちんとsourceなどで再起動を忘れずにしてください。  
# togglepx

togglepxは"toggle_proxy_automatically"という意味で、自動でプロキシを切り替えてくれるバイナリです。

---
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)
* [Addition](#addition)

---

# Features
このtogglepxバイナリは、ターミナルを起動させたときに、あなたの環境下のネットワークアドレスを判断し、
.gitconfigや.curlrcに記述されているProxy設定の行を自動でコメント/コメントアウトして、gitコマンドやcurlができるようにします。  

職場や大学がプロキシサーバーで自宅がそうでない場合、毎回.gitconfigなどのプロキシ行を書き換えなければならなかったため、そのような場合に使うといいと思います。  

---

## Installation
* **Goの開発環境があなたのPCに存在する場合**
```
go get github.com/ssabcire/togglepx
cd <GOPATH>/src/github.com/ssabcire/togglepx
go install
```
<br>

* **Goの開発環境があなたのPCに存在しない場合**  
Download a binary from [release page](https://github.com/ssabcire/togglepx/releases)

---

## Usage
1.シェルの設定ファイルに、togglepxバイナリのPATHを記述します。  
例  ~/.bashrc
```
~/go/src/togglepx/togglepx
``` 
  
2.次に、バイナリを作動させます。そうすることで、~/.togglepx/config.jsonが作成されます。  
例
```
source ~/.bashrc
```

3.~/.togglepx/config.jsonを編集します。
filepathには、プロキシ設定が書かれているファイルを記載してください。  
pxipには、プロキシサーバー下のネットワークアドレスを記載してください。  
例
```
{
  "filepath": "~/.gitconfig",
  "pxip": "192.168.1.0"
}
```

---

## Addition
このバイナリをさらに便利にするためのCLIも作成しています。  
togglepx/cmd/tpaにCLIのバイナリがあります。  
[CLIのREADME](https://github.com/ssabcire/togglepx/blob/master/cmd/README.md)

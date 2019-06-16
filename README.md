# togglepx

togglepxは"toggle_proxy_automatically"という意味で、自動でプロキシを切り替えてくれるバイナリです。

<br>
---
<br>

* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)
* [Example](#example)
* [Addition](#addition)

<br>

# Features
このtogglepxバイナリは、ターミナルを起動させたときに、あなたの環境下のネットワークアドレスを判断し、
.gitconfigや.curlrcに記述されているProxy設定の行を自動でコメント/コメントアウトして、gitコマンドやcurlができるようにします。  

職場や大学がプロキシサーバーで自宅がそうでない場合、毎回.gitconfigなどのプロキシ行を書き換えなければならなかったため、そのような場合に使うといいと思います。  

<br>

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

<br>

## Usage
#### Goの開発環境が存在する場合
1.バイナリを生成します。
```
go get github.com/ssabcire/togglepx
cd <GOPATH>/src/github.com/ssabcire/togglepx
go install
```
これで、\<GOPATH>/binにバイナリが配置されます。  
<br>
2.シェルの設定ファイルに、togglepxバイナリのPATHを記述します。  
```
<GOPATH>/bin/togglepx
```
<br>

3.次に、バイナリを作動させます。そうすることで、~/.togglepx/config.jsonが作成されます。
```
source ~/.bashrc
```
<br>

4.~/.togglepx/config.jsonを編集します。
filepathには、プロキシ設定が書かれているファイルを記載してください。  
pxipには、プロキシサーバー下のネットワークアドレスを記載してください。  
```
(例)
{
  "filepath": "~/.gitconfig",
  "pxip": "192.168.1.0"
}
```
<br><br>

#### Goの開発環境が存在しない場合
1.シェルの設定ファイルに、togglepxバイナリのPATHを記述します。  
~/.bashrc  
```
~/go/src/togglepx/togglepx
```
<br>

2.次に、バイナリを作動させます。そうすることで、~/.togglepx/config.jsonが作成されます。  
```
source ~/.bashrc
```

<br>

3.~/.togglepx/config.jsonを編集します。
filepathには、プロキシ設定が書かれているファイルを記載してください。  
pxipには、プロキシサーバー下のネットワークアドレスを記載してください。  
```
(例)
{
  "filepath": "~/.gitconfig",
  "pxip": "192.168.1.0"
}
```

<br>

## Example
もしプロキシサーバー下のネットワークアドレスが192.168.1.0で、.gitconfigのプロキシ設定を自動で切り替えてほしいとき  

```json:~/.togglepx/config.json  
{
  "filepath": "~/.gitconfig",
  "pxip": "192.168.1.0"
}
```  
<br>

次に、PATHを~/.bashrcに書き込みます。
```
~/.bashrc
```
<br><br>

こうすることで、プロキシサーバー下にいるときにシェルを起動すると、
```
# proxy=<プロキシサーバー>:<ポート>
```
の部分が
```
proxy=<プロキシサーバー>:<ポート>
```
に、自動でコメントアウトされます(プロキシサーバー下にいないときはコメントされます)。  
<br><br>

## Addition
このバイナリをさらに便利にするためのCLIも作成しています。  
togglepx/cmd/tpaにCLIのバイナリがあります。  
[CLIのREADME](https://github.com/ssabcire/togglepx/blob/master/cmd/README.md)

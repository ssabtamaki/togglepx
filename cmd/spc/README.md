# Switch Proxy CLI

sfpは、switchpxのバイナリをさらに使いやすくするためのCLIです

---
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)
    * [Quick Start](#quick-start)

---

## Features  
このCLIは、[ssabcire/switchpx](https://github.com/ssabcire/switchpx)のバイナリを補助するために作られたものです。  
switchpxバイナリは、ターミナルを起動したときに、あなたの環境下のネットワークアドレスを判断し、プロキシが記載された設定ファイルを自動でコメント/コメントアウトしてくれます。  
しかし、 そのswicthpxバイナリは、ターミナルが起動したときしかプロキシを切り替えてくれません。    
そこで、ターミナルを起動しているときにでも簡単に切り替えられるように、こちらのCLIがあります。

---

## Requirements
[Switchpxバイナリ](https://github.com/ssabcire/switchpx/releases)

---

## Installation
* **Goの開発環境があなたのPCに存在する場合**
```
go get github.com/ssabcire/spc
```
<br>

* **Goの開発環境があなたのPCに存在しない場合**  
Download a binary from [release page](https://github.com/ssabcire/spc/releases)
```
vim ~/.bashrc
<バイナリのPATHを記載>
```
 
 ---

## Usage
* -cancelip
* -cancelpath
* -checkpath
* -fpath <.gitconfigなどのプロキシが書かれた設定ファイル>
* -pxip <プロキシが設定されているネットワークアドレス>
* -switch

### Example
.gitconfigなどの設定ファイルを、自動コメント切り替えの対象にしたい場合
0. まずはじめに、switchpxバイナリを使用したことがないときは、1度だけ使用してください。(```~/.sfp/config.json```が作成されます)
1. 設定されているパスを確認します。```$ spc -checkpath``` と行うと、```<test>```が帰ってきます。
2. .gitconfigのPATHを設定します。```$ spc -fpath /Users/<ユーザー名>/.gitconfig```
3. 正しく設定されているか確認してみます。```$ spc -checkpath```   ```</Users/<ユーザー名>/.gitconfig>```
4. 設定されているアドレスを確認します。```$ spc -checkip``` を行うと、```<127.0.0.1>```が帰ってきます
5. ```$spc -pxip <ネットワークアドレス>```でプロキシが設定されているネットワークアドレスを選択します
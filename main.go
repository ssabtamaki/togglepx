package main

import (
	"fmt"
	"log"
	"os"
	"tpa/lib"
)

func main() {
	//ディレクトリの存在確認
	_, err := os.Stat(lib.JsonDir);
	if os.IsNotExist(err){
		err = os.MkdirAll(lib.JsonDir, 0777);
		if err != nil {
			log.Print("ディレクトリの作成に失敗しました", err)
			os.Exit(1)
		}
	}

	//設定ファイルから情報取得
	pathIPConfig := &lib.PathIPConfig{}
	err = pathIPConfig.ReadJsonTransfer(lib.JsonPath)
	if err != nil {
		fmt.Print("Jsonファイルから構造体への変換に失敗しました。")
		os.Exit(1)
	}

	//ネットワークアドレスの取得
	c := &lib.Client{Tst: &lib.Actual{}}
	netAddr, err := c.NetAddrPrint()
	if err != nil {
		log.Print("ネットワークアドレスの取得に失敗しました")
		os.Exit(1)
	}

	//プロキシ下のネットワークアドレスにいるとき
	if netAddr == pathIPConfig.FilePath {
		err = lib.SwitchProxyAuto(pathIPConfig.FilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "自動コメントアウトに失敗しました。", err)
		}
		os.Exit(0)
	}
	//プロキシ環境下以外のとき
	err = lib.SwitchProxyAuto(pathIPConfig.FilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "自動コメントアウトに失敗しました。", err)
	}
	os.Exit(0)
}

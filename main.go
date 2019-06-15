package main

import (
	"fmt"
	"log"
	"os"
	"tpa/lib"
)

func main() {
	pathIPConfig := &lib.PathIPConfig{}
	//設定ファイルから情報取得
	err := pathIPConfig.ReadJsonTransfer(lib.JsonPath)
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
}

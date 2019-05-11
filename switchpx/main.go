package main

import (
	"fmt"
	"log"
	"os"
	"switchpx/fproxy"
)

func main() {
	//ディレクトリの存在確認
	_, err := os.Stat(fproxy.JsonDir);
	if os.IsNotExist(err){
		err = os.MkdirAll(fproxy.JsonDir, 0777);
		if err != nil {
			log.Print("ディレクトリの作成に失敗しました", err)
			os.Exit(1)
		}
	}

	//設定ファイルから情報取得
	pathIPConfig := &fproxy.PathIPConfig{}
	err = pathIPConfig.ReadJsonTransfer()
	if err != nil {
		fmt.Print("Jsonファイルから構造体への変換に失敗しました。")
		os.Exit(1)
	}

	//ネットワークアドレスの取得
	c := &fproxy.Client{Tst: &fproxy.Actual{}}
	netAddr, err := c.NetAddrPrint()
	if err != nil {
		log.Print("ネットワークアドレスの取得に失敗しました")
		os.Exit(1)
	}

	//プロキシ下のネットワークアドレスにいるとき
	if netAddr == pathIPConfig.FilePath {
		err = fproxy.SwitchProxyAuto(pathIPConfig.FilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "自動コメントアウトに失敗しました。", err)
		}
		os.Exit(0)
	}
	//プロキシ環境下以外のとき
	err = fproxy.SwitchProxyAuto(pathIPConfig.FilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "自動コメントアウトに失敗しました。", err)
	}
	os.Exit(0)
}

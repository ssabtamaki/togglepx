package main

import (
	"os"
	"stepupgo/cli"
)

func main() {
	//構造体を作る
	p := cli.PathIPConfig{}
	//ファイル読み込んで構造体に適用
	err := p.ReadJsonTransfer(cli.JsonPath)
	if err != nil {
		os.Exit(1)
	}
	//cli起動、cliに構造体渡す
	stream := cli.Stream{os.Stdout, os.Stderr}
	os.Exit(stream.Run(os.Args, &p))
	//終了
}

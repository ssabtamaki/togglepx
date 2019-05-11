package main

import (
	"os"
	"spc/cli"
)

func main() {
	p := cli.PathIPConfig{}
	//ファイルを読み込んで構造体に適用
	err := p.ReadJsonTransfer(cli.JsonPath)
	if err != nil {
		os.Exit(1)
	}
	stream := cli.Stream{os.Stdout, os.Stderr}
	//cli起動、cliに構造体渡す
	os.Exit(stream.Run(os.Args, &p))
}

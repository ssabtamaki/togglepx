package main

import (
	"fmt"
	"log"
	"os"
	"stepupgo/cli"
	"stepupgo/fproxy"
)

func main() {
	if (cli.Fpath == "") || (cli.PxIP.String() == "") {
		fmt.Fprintln(os.Stderr, "IPアドレスもしくは対象ファイルが設定	されていません。プロキシの自動切り替えは行わず、CLIを起動します")
		stream := &cli.Stream{OutStream: os.Stdout, ErrStream: os.Stderr}
		os.Exit(stream.Run(os.Args))
	}

	netIPv4, err := fproxy.GetNetIPv4()
	if err != nil {
		log.Println("現在の環境化のネットワークアドレスの取得に失敗しました。プロキシの自動切り替えは行わず、CLIを起動します")
		stream := &cli.Stream{OutStream: os.Stdout, ErrStream: os.Stderr}
		os.Exit(stream.Run(os.Args))
	}

	//プロキシ下のネットワークアドレスにいるとき
	if netIPv4.String() == cli.PxIP.String() {
		err = fproxy.SwitchProxyAuto(cli.Fpath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "自動コメントアウトに失敗しました。プロキシの自動切り替えは行わず、CLIを起動します")
		}
		stream := &cli.Stream{OutStream: os.Stdout, ErrStream: os.Stderr}
		os.Exit(stream.Run(os.Args))
	}
	//プロキシ環境下以外のとき
	err = fproxy.SwitchProxyAuto(cli.Fpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "自動コメントアウトに失敗しました。プロキシの自動切り替えは行わず、CLIを起動します")
	}
	stream := &cli.Stream{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(stream.Run(os.Args))
}
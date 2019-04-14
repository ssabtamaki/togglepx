package main

import (
	"fmt"
	"log"
	"os"
	"stepupgo/cli"
	"stepupgo/fproxy"
)

//mainの実行順序はどうしたほうがベストなのか？
func main() {
	if !((cli.Fpath == "") && (cli.PxIP == "")) {
		fmt.Fprintln(os.Stderr, "IPアドレスもしくは対象ファイルが設定されていません。プロキシの自動切り替えは行わず、CLIを起動します")
		stream := &cli.Stream{OutStream: os.Stdout, ErrStream: os.Stderr}
		os.Exit(stream.Run(os.Args))
	}

	netIPv4, err := fproxy.GetNetIPv4()
	if err != nil {
		log.Println("Error to netIPv4")
		os.Exit(1)
	}
	//プロキシ下のネットワークアドレスにいるとき
	if netIPv4.String() == cli.PxIP {
		err = fproxy.SwitchProxyAuto(cli.Fpath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to Comment Out.　自動コメントアウトに失敗しました")
		}
		return		//os.Exit(0)のほうがいい？
	}
	//プロキシ環境下以外のとき
	err = fproxy.SwitchProxyAuto(cli.Fpath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to Comment Out. 自動コメントアウトに失敗しました")
	}
	stream := &cli.Stream{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(stream.Run(os.Args))
}

package main

import (
	"fmt"
	"log"
	"os"
	"stepupgo/fproxy"
)

const (
	gitConfig = "/Users/ssab/.gitconfig.org"
	//プロキシ下にある大学のネットワークアドレス.
	univIP = "192.168.16.0"
	FilePath = "/test/cli"
)

func main() {
	netIPv4, err := fproxy.GetNetIPv4()
	if err != nil {
		log.Println("Error to netIPv4")
		os.Exit(1)
	}

	//大学にいるとき
	if netIPv4.String() == univIP {
		err = fproxy.SwitchProxyAuto(gitConfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to Comment Out")
		}
		return
	}
	//大学以外のとき
	err = fproxy.SwitchProxyAuto(gitConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to Comment Out")
	}
}

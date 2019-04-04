package main

import (
	"fmt"
	"os"
	"stepupgo/fproxy"
)

const (
	gitConfig = "/Users/ssab/.gitconfig.org"
	//プロキシ下にある大学のネットワークアドレス.
	univIP = "192.168.16.0"
)

func main() {
	netIPv4, err := fproxy.GetNetIPv4()
	if err != nil {
		//log.Fatal("Error to netIPv4")
		fmt.Fprintln(os.Stderr, "Error to netIPv4")
		os.Exit(1)
	}

	//大学にいるとき
	if netIPv4.String() == univIP {
		err = fproxy.ProxySubComment(gitConfig)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to Comment Out")
		}
		return
	}
	//大学以外のとき
	err = fproxy.ProxyAddComment(gitConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to Comment Out")
	}
}

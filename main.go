package main

import (
	"log"
	"stepupgo/fproxy"
)

const (
	gitConfig = "/Users/ssab/.gitConfig.org"
	//プロキシ下にある大学のネットワークアドレス.
	univIP = "192.168.16.0"
	//test = "proxy.txt"
)

func main() {
	netIPv4, err := fproxy.GetNetIPv4()
	if err != nil {
		log.Fatal("Error to netIPv4")
	}

	//大学以外のとき
	if netIPv4.String() != univIP {
		err = fproxy.ProxySubComment(gitConfig)
		if err != nil {
			fmt.Println("Failed to Comment Out")
		}
	}
	//大学にいるとき
	err = fproxy.ProxyAddComment(gitConfig)
	if err != nil {
		fmt.Println("Failed to Comment Out")
	}
}
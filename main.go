package main

import (
	"fmt"
	"log"
)

const (
	gitConfig = "/Users/ssab/.gitConfig.org"
	//test = "proxy.txt"
	//プロキシ下にある大学のネットワークアドレス.
	univIP = "192.168.16.0"
)

func main() {
	netIPv4, err := getNetIPv4()
	if err != nil {
		log.Fatal("Error to netIPv4")
	}

	//大学以外のとき
	if netIPv4.String() != univIP {
		err = proxySubComment(gitConfig)
		if err != nil {
			fmt.Println("Failed to Comment Out")
		}
	}
	//大学にいるとき
	err = proxyAddComment(gitConfig)
	if err != nil {
		fmt.Println("Failed to Comment Out")
	}
}

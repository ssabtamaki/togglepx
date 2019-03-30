package main

import (
	"fmt"
	"log"
)

func main() {
	netIPv4, err := getNetIPv4()
	if err != nil {
		log.Fatal("Error to netIPv4")
	}

	//大学以外のとき
	if netIPv4.String() != "192.168.16.0" {
		err = proxySubComment("proxy.txt")
		if err != nil {
			fmt.Println("Failed to Comment Out")
		}
	}
	//大学にいるとき
	err = proxyAddComment("proxy.txt")
	if err != nil {
		fmt.Println("Failed to Comment Out")
	}
}

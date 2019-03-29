package main

import (
	"log"
	"fmt"
	"os"
)

func main() {
	netIPv4 ,err := getNetIPv4()
	if err != nil {
		log.Fatal("Error to netIPv4")
	}

	file, err := os.OpenFile("proxy.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error to Open file")
	}
	defer file.Close()

	//大学以外のとき
	if netIPv4.String() != "192.168.16.0" {
		err = proxySubComment(file)
		if err != nil {
			fmt.Println("Failed to Comment Out")
		}
	}
	//大学にいるとき
	err = proxyAddComment(file)
	if err != nil {
		fmt.Println("Failed to Comment Out")
	}
}

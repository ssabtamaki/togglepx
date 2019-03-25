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

	file, err := os.OpenFile("proxy.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error to Open file")
	}
	defer file.Close()

	if netIPv4.String() != "192.168.11.0" {
		err = proxyAddComment(file)
		if err != nil {
			fmt.Fprint(stderr, "Failed to Comment Out")
		}
	}
	err = proxySubComment(file)
	if err != nil {
		fmt.Fprint(stderr, "Failed to Comment Out")
	}
}

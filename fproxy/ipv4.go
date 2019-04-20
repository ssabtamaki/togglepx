package fproxy

import (
	"fmt"
	"net"
)

//ネットワークアドレスを取得する
func GetNetIPv4() (netIPv4 net.IP, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("cannot get address.")
		return
	}

	var ipv4 net.IP
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipv4 = ipnet.IP
			}
		}
	}

	mask := ipv4.DefaultMask()
	netIPv4 = ipv4.Mask(mask)
	return
}

// IPv4アドレスを求めるプログラム
package fproxy

import (
	"fmt"
	"net"
)

//ネットワークアドレスを取得する
func GetNetIPv4() (netIPv4 net.IP, err error) {
	// すべてのアドレスを取得
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("cannot get address.")
		return
	}

	var ipv4 net.IP
	// IPv4アドレスを取得
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		// ipnetが真で、ループバックアドレスを除外
		if ok && !ipnet.IP.IsLoopback() {
			//IPv4に変換できるとき
			if ipnet.IP.To4() != nil {
				ipv4 = ipnet.IP
			}
		}
	}

	// ipv4からマスク値を求める
	mask := ipv4.DefaultMask()
	netIPv4 = ipv4.Mask(mask)
	return
}

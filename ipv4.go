package main

import (
	"fmt"
	"net"
)

type Tster interface {
	GetNetAddr() (net.IP, error)
}

type Client struct {
	Tst Tster
}

type Actual struct{}

func (a *Actual) GetNetAddr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("cannot get address.")
		return nil, err
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
	netAddr := ipv4.Mask(mask)
	return netAddr, err
}

//テスト用メソッド
func (c *Client) NetAddrPrint() (string, error) {
	netAddr, err := c.Tst.GetNetAddr()
	if err != nil {
		return "", err
	}
	return netAddr.String(), nil
}

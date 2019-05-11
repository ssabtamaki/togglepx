package test

import (
	"net"
	"switchpx/fproxy"
	"testing"
)

type dummyMock struct{}

//ダミーのネットワークアドレスを返すように、モックを実装
func (d *dummyMock) GetNetAddr() (net.IP, error) {
	return net.IPv4(127, 0, 0, 0), nil
}

func TestNetIP(t *testing.T) {
	c := &fproxy.Client{Tst: &dummyMock{}}
	netAddr, err := c.NetAddrPrint()
	if err != nil {
		t.Errorf("caused error:%s", err)
	}
	if expected := "127.0.0.0"; expected != netAddr {
		t.Errorf("want %s, got %s", expected, netAddr)
	}

	c = &fproxy.Client{Tst: &fproxy.Actual{}}
	netAddr, err = c.NetAddrPrint()
	if err != nil {
		t.Errorf("caused error:%s", err)
	}
	if expected := "192.168.20.0"; expected != netAddr {
		t.Errorf("want %s, got %s", expected, netAddr)
	}
}

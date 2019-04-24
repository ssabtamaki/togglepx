package test

import (
	"net"
	"stepupgo/fproxy"
	"testing"
)

type dummyTster struct { /*Tsterが暗黙的に宣言されてる?*/
}

//ダミーのネットワークアドレスを返すように、モックを実装
func (d *dummyTster) GetNetAddr() (net.IP, error) {
	return net.IPv4(127, 0, 0, 0), nil
}

func TestNetIP(t *testing.T) {
	c := &fproxy.Client{Tst: &dummyTster{}}
	netAddr, err := c.NetPrint()
	if err != nil {
		t.Errorf("caused error:%s", err)
	}
	if expected := "127.0.0.0"; expected != netAddr {
		t.Errorf("want %s, got %s", expected, netAddr)
	}

	c = &fproxy.Client{Tst: &fproxy.Actual{}}
	netAddr, err = c.NetPrint()
	if err != nil {
		t.Errorf("caused error:%s", err)
	}
	if expected := "192.168.20.0"; expected != netAddr {
		t.Errorf("want %s, got %s", expected, netAddr)
	}
}

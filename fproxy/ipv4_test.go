package fproxy

import (
	"testing"
)

const univIP = "192.168.16.0"

func Test_getNetIPv4(t *testing.T) {
	netIPv4, err := GetNetIPv4()
	if err != nil {
		t.Errorf("error to getNetIPv4")
	}
	if netIPv4.String() != univIP {
		t.Errorf("Error to get NetworkIPaddress")
	}
}

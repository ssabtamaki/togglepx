package main

import (
	"testing"
	)

func Test_getNetIPv4(t *testing.T) {
	netIPv4 ,err := getNetIPv4()
	if err != nil {
		t.Errorf("error to getNetIPv4")
	}
	if netIPv4.String() != "192.168.11.0" {
		t.Errorf("Error to get NetworkIPaddress")
	}
}
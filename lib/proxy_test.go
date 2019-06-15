package lib_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"togglepx/lib"
)

func Test_SwitchProxyAuto(t *testing.T) {
	filename := "proxy.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Error("Error to Open File", err)
	}
	defer file.Close()

	_, err = file.Write([]byte("proxy=<>:<>"))
	if err != nil {
		t.Error("Error to Write to File")
	}

	err = lib.ToggleProxyAuto(filename)
	if err != nil {
		t.Errorf("Error SwitchProxyAdd")
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Error to Read File")
	}

	if !strings.Contains(string(input), lib.ExportPx) {
		t.Error("Error to Proxy Replace")
	}

	err = lib.ToggleProxyAuto(filename)
	if err != nil {
		t.Errorf("Error SwitchProxySub")
	}
	input, err = ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Error to Read File")
	}

	if strings.Contains(string(input), lib.ExportCommentPx) {
		t.Error("Error to Proxy Replace")
	}
}

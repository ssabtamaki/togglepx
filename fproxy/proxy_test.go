package fproxy

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_SwitchProxyAuto(t *testing.T) {
	//ファイルがあったら上書き、なかったら新規作成。最後に追記ではない
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

	err = SwitchProxyAuto(filename)
	if err != nil {
		t.Errorf("Error SwitchProxyAdd")
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Error to Read File")
	}

	if !strings.Contains(string(input), cPx) {
		t.Error("Error to Proxy Replace")
	}

	err = SwitchProxyAuto(filename)
	if err != nil {
		t.Errorf("Error SwitchProxySub")
	}
	input, err = ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Error to Read File")
	}

	if strings.Contains(string(input), cPx) {
		t.Error("Error to Proxy Replace")
	}
}

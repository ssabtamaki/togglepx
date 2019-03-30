package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_proxyAddComment(t *testing.T) {
	//ファイルがあったら上書き、なかったら新規作成。最後に追記ではない
	filename := "proxy.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Error("Error to Open File", err)
	}
	defer file.Close()

	_, err = file.Write([]byte("proxy=kanazawa-it.ac.jp:8080"))
	if err != nil {
		t.Error("Error to Write to File")
	}

	err = proxyAddComment(filename)
	if err != nil {
		t.Errorf("Error ProxyAddComment")
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Error to Read File")
	}

	if !strings.Contains(string(input), commentProxy) {
		t.Error("Error to Proxy Replace")
	}
}

func Test_proxySubComment(t *testing.T) {
	//ファイルがあったら上書き、なかったら新規作成。最後に追記ではない
	filename := "proxy.txt"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Error("Error to Open File", err)
	}
	defer file.Close()

	_, err = file.Write([]byte("# proxy=kanazawa-it.ac.jp:8080"))
	if err != nil {
		t.Error("Error to Write to File")
	}

	err = proxySubComment(filename)
	if err != nil {
		t.Errorf("Error ProxySubComment")
	}

	input, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error("Error to Read File")
	}
	//ここを正規表現に変えたい
	if strings.Contains(string(input), commentProxy) {
		t.Error("Error to Proxy Replace")
	}
}

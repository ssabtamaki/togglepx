package test

import (
	"fmt"
	"os"
	"spc/cli"
	"switchpx/fproxy"
	"testing"
)

func init() {
	//ファイルがあったら削除させる
	_, err := os.Stat(fproxy.JsonPath)
	if os.IsExist(err) {
		err := os.Remove(fproxy.JsonPath)
		if err != nil {
			fmt.Println("ファイルの削除に失敗", err)
			os.Exit(1)
		}
		fmt.Println("テストのため、設定ファイルを削除しました")
	}
}

func Test_newCreateJsonFile(t *testing.T) {
	err := cli.NewCreateJsonFile()
	if err != nil {
		t.Errorf("ファイル作成エラー")
	}
	p := cli.PathIPConfig{}
	err = p.ReadJsonTransfer(cli.JsonPath)
	if err != nil {
		t.Errorf("transfer err")
	}
	expected := "test"
	if expected != p.FilePath {
		t.Errorf("get %v, want %v", p.FilePath, expected)
	}
	expected = "127.0.0.1"
	if expected != p.PxIP {
		t.Errorf("get %v, want %v", p.PxIP, expected)
	}
}

func Test_WriteToJsonFile(t *testing.T) {
	p := cli.PathIPConfig{"abc", "192.168.1.1"}
	err := cli.WriteToJsonFile(&p)
	if err != nil {
		t.Errorf("ファイルに書き込みに失敗")
	}
	q := cli.PathIPConfig{}
	err = q.ReadJsonTransfer(cli.JsonPath)
	if err != nil {
		t.Errorf("ReadJsonTransger失敗")
	}
	expected := "abc"
	if expected != q.FilePath {
		t.Errorf("get %v, want %v", p.FilePath, expected)
	}
	expected = "192.168.1.1"
	if expected != q.PxIP {
		t.Errorf("get %v, want %v", p.FilePath, expected)
	}
}

package test

import (
	"fmt"
	"log"
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


//package test
//
//import (
//"fmt"
//"log"
//"os"
//"switchpx/fproxy"
//"testing"
//)
//
//func init() {
//	//ディレクトリの存在確認、なければ作成
//	_, err := os.Stat(fproxy.JsonDir);
//	if os.IsNotExist(err){
//		err = os.MkdirAll(fproxy.JsonDir, 0777);
//		if err != nil {
//			log.Print("ディレクトリの作成に失敗しました", err)
//			os.Exit(1)
//		}
//	}
//
//	//ファイルがあったら削除させる
//	_, err = os.Stat(fproxy.JsonPath)
//	if os.IsExist(err) {
//		err := os.Remove(fproxy.JsonPath)
//		if err != nil {
//			fmt.Println("ファイルの削除に失敗", err)
//			os.Exit(1)
//		}
//		fmt.Println("テストのため、設定ファイルを削除しました")
//	}
//}
//
//func Test_config(t *testing.T) {
//	pathIPConfig := &fproxy.PathIPConfig{}
//	err := pathIPConfig.ReadJsonTransfer()
//	if err != nil {
//		t.Errorf("ReadJsonTransfer()でエラーが発生しました. %v", err)
//	}
//
//	//ファイルの存在確認
//	_, err = os.Stat(fproxy.JsonPath)
//	if os.IsNotExist(err) {
//		t.Errorf("設定ファイルの作成に失敗しています。本来なら作成されます")
//	}
//
//	expected := "test"
//	if expected != pathIPConfig.FilePath {
//		t.Errorf("want:%v, got:%v", expected, pathIPConfig.FilePath)
//	}
//
//	expected = "127.0.0.1"
//	if expected != pathIPConfig.PxIP {
//		t.Errorf("want:%s, got:%v", expected, pathIPConfig.PxIP)
//	}
//}

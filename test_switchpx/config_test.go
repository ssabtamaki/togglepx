package test

import (
	"fmt"
	"log"
	"os"
	"switchpx/fproxy"
	"testing"
)

func init() {
	//ディレクトリの存在確認、なければ作成
	_, err := os.Stat(fproxy.JsonDir);
	if os.IsNotExist(err){
		err = os.MkdirAll(fproxy.JsonDir, 0777);
		if err != nil {
			log.Print("ディレクトリの作成に失敗しました", err)
			os.Exit(1)
		}
	}

	//ファイルがあったら削除させる
	_, err = os.Stat(fproxy.JsonPath)
	if os.IsExist(err) {
		err := os.Remove(fproxy.JsonPath)
		if err != nil {
			fmt.Println("ファイルの削除に失敗", err)
			os.Exit(1)
		}
		fmt.Println("テストのため、設定ファイルを削除しました")
	}
}

func Test_config(t *testing.T) {
	pathIPConfig := &fproxy.PathIPConfig{}
	err := pathIPConfig.ReadJsonTransfer()
	if err != nil {
		t.Errorf("ReadJsonTransfer()でエラーが発生しました. %v", err)
	}

	//ファイルの存在確認
	_, err = os.Stat(fproxy.JsonPath)
	if os.IsNotExist(err) {
		t.Errorf("設定ファイルの作成に失敗しています。本来なら作成されます")
	}

	expected := "test"
	if expected != pathIPConfig.FilePath {
		t.Errorf("want:%v, got:%v", expected, pathIPConfig.FilePath)
	}

	expected = "127.0.0.1"
	if expected != pathIPConfig.PxIP {
		t.Errorf("want:%s, got:%v", expected, pathIPConfig.PxIP)
	}
}

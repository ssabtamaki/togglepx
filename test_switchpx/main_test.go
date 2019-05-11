// +build integration

package test

import (
	"switchpx/fproxy"
	"testing"
)

func Test_main(t *testing.T) {
	_, err := os.Stat(fproxy.JsonDir);
	if os.IsNotExist(err){
		err = os.MkdirAll(fproxy.JsonDir, 0777);
		if err != nil {
			t.Errorf("ディレクトリの作成に失敗しました", err)
		}
	}

	//設定ファイルからの取得。ファイルがないときは作成する
	pathIPConfig := &fproxy.PathIPConfig{}
	err = pathIPConfig.ReadJsonTransfer()
	if err != nil {
		t.Errorf("ReadJsonTransfer()でエラーが発生しました. %v", err)
	}
	expected := "test"
	if expected != pathIPConfig.FilePath {
		t.Errorf("want:%v, got:%v", expected, pathIPConfig.FilePath)
	}
	expected = "127.0.0.1"
	if expected != pathIPConfig.PxIP {
		t.Errorf("want:%s, got:%v", expected, pathIPConfig.PxIP)
	}

	//ネットワークアドレスの取得
	c := &fproxy.Client{Tst: &dummyMock{}}
	netAddr, err := c.NetAddrPrint()
	if err != nil {
		t.Errorf("caused error:%s", err)
	}
	if expected := "127.0.0.0"; expected != netAddr {
		t.Errorf("want %s, got %s", expected, netAddr)
	}

	//テスト用のネットワークアドレスと設定ファイルのアドレスの一致確認
	if netAddr == pathIPConfig.PxIP {
		err = fproxy.SwitchProxyAuto(pathIPConfig.FilePath)
		if err != nil {
			t.Errorf("自動コメントアウトに失敗しました。")
		}
	}
}

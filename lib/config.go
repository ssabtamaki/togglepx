package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type PathIPConfig struct {
	FilePath string `json:"filepath"`
	PxIP     string `json:"pxip"`
}

var jsonDir = func() string {
	//現在のユーザーを取得
	user, err := user.Current()
	if err != nil {
		return ""
	}
	jsonPath := filepath.Join(user.HomeDir, ".tpa")
	return jsonPath
}()

var JsonPath = func() string {
	jsonPath := filepath.Join(jsonDir, "config.json")
	return jsonPath
}()

//ディレクトリの存在確認をして、なければ作成をする
func ifDirNotExistMkdir() error {
	_, err := os.Stat(jsonDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(jsonDir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func ifFileNotExistMkfile() error {
	_, err := os.Stat(JsonPath)
	if os.IsNotExist(err) {
		err = newCreateJsonFile()
		if err != nil {
			return err
		}
	}
	return nil
}

//JSONファイルの作成
func newCreateJsonFile() error {
	content := []byte(`
{
  "filepath": "test",
  "pxip": "127.0.0.1"
}
	`)
	err := ioutil.WriteFile(JsonPath, content, 0666)
	if err != nil {
		return err
	}
	return nil
}

//jsonを読みこみ、構造体pに渡す
func (p *PathIPConfig) ReadJsonTransfer(JsonPath string) error {
	err := ifDirNotExistMkdir()
	if err != nil {
		return err
	}
	err = ifFileNotExistMkfile()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(JsonPath)
	if err != nil {
		fmt.Println("ファイルの読み込みに失敗", err)
		return err
	}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	return nil
}

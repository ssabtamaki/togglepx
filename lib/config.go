package cli

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

var JsonPath = func() string {
	//現在のユーザーを取得
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return filepath.Join(user.HomeDir, ".sfp", "config.json")
}()

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
	//ファイルの存在確認
	_, err := os.Stat(JsonPath)
	if os.IsNotExist(err) {
		err = newCreateJsonFile()
		if err != nil {
			return err
		}
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

func writeToJsonFile(p *PathIPConfig) error {
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(JsonPath, json, 0666)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"testing"
)

func Test_proxyAddComment(t *testing.T) {
	//ファイルがあったら上書き、なかったら新規作成。最後に追記ではない
	file, err := os.OpenFile("proxy.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Error("Error to Open File", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "kanazawa-it.ac.jp:8080")

	err := ProxyAddComment(file)
	if err != nil {
		t.Errorf("Error ProxyAddComment")
	}

	//以下で、正規表現をして、プロキシが書かれている行にコメントアウトができているかをテストする
	scanner := bufio.Newcanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//プロキシが書かれている行を見つけられるかを調べる
		rex := regexp.MustCompile(`kanazawa-it.ac.jp:8080`)
		//マッチ失敗時
		if !rex.MatchString(line) {
			t.Errorf("Error to regexp proxy")
			return
		}
		//プロキシが記載されている行の先頭に、コメントアウトがされているかを調べる
		coRex := regexp.MustCompile(`# kanazawa-it.ac.jp:8080`)
		//マッチ失敗時
		if !coRex.MatchString(line) {
			t.Error("Error to write # to file", err)
		}
	}
}

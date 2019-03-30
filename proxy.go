// プロキシファイルに#を加える関数と、抜き取る関数
package main

import (
	"io/ioutil"
	"strings"
)

func proxyAddComment(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//ここの部分、正規表現で完全一致に切るように変更したほうがいいかもしれない
	//変更の余地あり
	//もうすでにコメントアウトされているときは何もしない
	if !strings.Contains(string(input), "# kanazawa-it.ac.jp") {
		output := strings.Replace(string(input), "kanazawa-it.ac.jp", "# kanazawa-it.ac.jp", 1)
		err = ioutil.WriteFile(filename, []byte(output), 0666)
	}
	if err != nil {
		return
	}
	return
}

func proxySubComment(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//ここの部分、正規表現で完全一致に切るように変更したほうがいいかもしれない
	//変更の余地あり
	if strings.Contains(string(input), "kanazawa-it.ac.jp") {
		output := strings.Replace(string(input), "# kanazawa-it.ac.jp", "kanazawa-it.ac.jp", 1)
		err = ioutil.WriteFile(filename, []byte(output), 0666)
	}
	if err != nil {
		return
	}
	return
}

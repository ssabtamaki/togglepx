// プロキシファイルに#を加える関数と、抜き取る関数
package fproxy

import (
	"io/ioutil"
	"strings"
)

const (
	noCommentProxy = "proxy="
	CommentProxy   = "# proxy="
)

//filenameのPROXY行の先頭に#を追加する
func ProxyAddComment(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//"proxy"があるが、 "# proxy"ではないとき
	if strings.Contains(string(input), noCommentProxy) && !strings.Contains(string(input), CommentProxy) {
		output := strings.Replace(string(input), noCommentProxy, CommentProxy, 1)
		err = ioutil.WriteFile(filename, []byte(output), 0666)
	}
	if err != nil {
		return
	}
	return
}

//filenameのPROXYの先頭に#を削除する
func ProxySubComment(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if strings.Contains(string(input), CommentProxy) {
		output := strings.Replace(string(input), CommentProxy, noCommentProxy, 1)
		err = ioutil.WriteFile(filename, []byte(output), 0666)
	}
	if err != nil {
		return
	}
	return
}

// プロキシファイルに#を加える関数と、抜き取る関数
package fproxy

import (
	"io/ioutil"
	"strings"
)

const (
	px  = "proxy="
	cPx = "# proxy="
)

//filenameのPROXY行の先頭に#を追加する
func SwitchProxyAuto(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	output := strings.NewReplacer(px, cPx, cPx, px).Replace(string(input))
	err = ioutil.WriteFile(filename, []byte(output), 0666)
	if err != nil {
		return
	}
	return
}

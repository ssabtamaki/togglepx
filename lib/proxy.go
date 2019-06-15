package lib

import (
	"io/ioutil"
	"strings"
)

const (
	px  = "proxy="
	commentPx = "# proxy="
)

//ファイルのプロキシが書かれている行に#を入れたり抜いたりする
func ToggleProxyAuto(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	output := strings.NewReplacer(px, commentPx, commentPx, px).Replace(string(input))
	err = ioutil.WriteFile(filename, []byte(output), 0666)
	if err != nil {
		return err
	}
	return nil
}

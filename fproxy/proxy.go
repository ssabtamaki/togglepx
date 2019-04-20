package fproxy

import (
	"io/ioutil"
	"strings"
)

const (
	Px  = "proxy="
	Cpx = "# proxy="
)

//ファイルのプロキシが書かれている行に#を入れたり抜いたりする
func SwitchProxyAuto(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	output := strings.NewReplacer(Px, Cpx, Cpx, Px).Replace(string(input))
	err = ioutil.WriteFile(filename, []byte(output), 0666)
	if err != nil {
		return err
	}
	return nil
}

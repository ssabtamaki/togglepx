package lib

import (
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	px        = "proxy="
	commentPx = "# proxy="
)

func ProxyOFF(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//output := strings.NewReplacer(px, commentPx).Replace(string(input))
	output := input
	if regexp.MustCompile(px).MatchString(string(input)) && !regexp.MustCompile(commentPx).MatchString(string(input)) {
		//Replaceする
		output = []byte(strings.NewReplacer(px, commentPx).Replace(string(input)))
	}
	err = ioutil.WriteFile(filename, output, 0666)
	if err != nil {
		return err
	}
	return nil
}

func ProxyON(filename string) (err error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	output := input
	if regexp.MustCompile(commentPx).MatchString(string(input)) {
		output = []byte(strings.NewReplacer(commentPx, px).Replace(string(input)))
	}
	//output := strings.NewReplacer(commentPx, px).Replace(string(input))
	err = ioutil.WriteFile(filename, output, 0666)
	if err != nil {
		return err
	}
	return nil
}

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

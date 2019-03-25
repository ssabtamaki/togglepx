// プロキシファイルに#を加える関数と、抜き取る関数
package main

import (
	"bufio"
	"os"
	"regexp"
)

// プロキシが書かれた行の先頭に#,コメントアウトを入れる(ファイルを書き換える)
func proxyAddComment(file *os.File) (err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rep := regexp.MustCompile(`kanazawa-it.ac.jp`)
		if rep.MatchString(line) {
			line = rep.ReplaceAllString(line, "# kanazawa-it.ac.jp:8080")
			//以下は、lineの行を書き換えたい
			//~~~
		}
	}
	return
}

// プロキシが書かれた行の先頭の#を抜く（実際には書き換える）
func proxySubComment(file *os.File) (err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rep := regexp.MustCompile(`kanazawa-it.ac.jp`)
		if rep.MatchString(line) {
			line = rep.ReplaceAllString(line, "kanazawa-it.ac.jp:8080")
			//以下で、ファイルの内容を書き換えたい
			//~~~
		}
	}
	return
}

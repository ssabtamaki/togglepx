package cli

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type cli struct {
	outStream, errStream io.Writer
}

//var fpath string

func (c *cli) Run(args []string) int {
	//switch file proxy = swfpx
	flags := flag.NewFlagSet("swfpx", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	var fpath string	//ここ少し問題
	//引数でパスが入力されたら、fpathに代入される
	flags.StringVar(&fpath, "fpath", fpath, "Specify the path of the file you want to target proxy on/off. プロキシのオンオフ対象のファイルのPATHを設定します")
	var cancel bool
	flags.BoolVar(&cancel, "cancel", false, "Cancel the registered path. 登録されているパスを解除します")
	var check bool
	flags.BoolVar(&check, "check", false, "Check the currently set path. 現在設定されているパスを確認します")
	var effective bool
	var ineffective bool
	flags.BoolVar(&effective, "effective", false, "Uncomment the target file and activate the proxy. 対象ファイルのコメントをはずし、プロキシを有効化します")
	flags.BoolVar(&ineffective, "ineffective", false, "Comment the target file and disable the proxy. 対象ファイルにコメントをつけ、プロキシを無効化します")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if fpath != "" {
		fmt.Fprintf(c.errStream, "プロキシのオンオフ対象のファイルを、%sにPATHを設定しました", fpath)
	}
	if cancel {
		fmt.Fprint(c.errStream, "設定されているPATHを取り消しました")
	}
	if check {
		fmt.Fprintf(c.errStream, "現在設定されているPATHは%sです", fpath)
	}
	if effective {
		fmt.Fprint(c.errStream, "対象ファイルのコメントをはずし、プロキシを有効化します")
	}
	if ineffective {
		fmt.Fprint(c.errStream, "対象ファイルにコメントをつけ、プロキシを無効化しました")
	}
	return ExitCodeOK
}
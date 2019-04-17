package cli

import (
	"flag"
	"fmt"
	"io"
	"stepupgo/fproxy"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeExeFlagError
)

type Stream struct {
	OutStream, ErrStream io.Writer
}

var Fpath string
var PxIP string

//CLIのコマンドの設定と実行
func (s *Stream) Run(args []string) int {
	//swfpx = switch file proxy
	flags := flag.NewFlagSet("swfpx", flag.ContinueOnError)
	flags.SetOutput(s.ErrStream)

	//ここ、StringVarじゃなくてIP型で、flags.Varでやるのもあり？
	//そうしたほうが、IPアドレスじゃなかったときにエラーが吐ける
	//エラーを吐くときはどうする？
	flags.StringVar(&PxIP, "pxip", PxIP, "Register a network address. When registering, please register the address under proxy environment. ネットワークアドレスを登録します。登録するときは、プロキシ環境下のアドレスを登録してください\n")
	var checkIP bool
	flags.BoolVar(&checkIP, "checkip", false, "Check the registered network address value. 登録したネットワークアドレスの値を確認します\n")
	var cancelIP bool
	flags.BoolVar(&cancelIP, "cancelip", false, "Cancel the registered network address value. 登録したネットワークアドレスの値を解除します\n")

	flags.StringVar(&Fpath, "filepath", Fpath, "Set the PATH of the target proxy file on/off. プロキシのオンオフ対象のファイルのPATHを設定します\n")
	var cancelPath bool
	flags.BoolVar(&cancelPath, "cancelpath", false, "Cancel the registered path. 登録されているパスを解除します\n")
	var checkPath bool
	flags.BoolVar(&checkPath, "checkpath", false, "Check the currently set path. 現在設定されているパスを確認します\n")

	var effective bool
	var ineffective bool
	flags.BoolVar(&effective, "effective", false, "Uncomment the target file and activate the proxy. 対象ファイルのコメントをはずし、プロキシを有効化します\n")
	flags.BoolVar(&ineffective, "ineffective", false, "Comment the target file and disable the proxy. 対象ファイルにコメントをつけ、プロキシを無効化します\n")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if PxIP != "" {
		fmt.Fprintf(s.ErrStream, "ネットワークアドレス%sを登録しました\n", PxIP)
	}
	if checkIP {
		fmt.Fprintf(s.ErrStream, "現在設定されているネットワークアドレスは%sです\n", PxIP)
	}
	if cancelIP {
		fmt.Fprint(s.ErrStream, "設定されているネットワークアドレスを取り消しました\n")
	}
	if Fpath != "" {
		fmt.Fprintf(s.ErrStream, "プロキシのオンオフ対象のファイルを、%sにPATHを設定しました\n", Fpath)
	}
	if cancelPath {
		fmt.Fprint(s.ErrStream, "設定されているPATHを取り消しました\n")
	}
	if checkPath {
		fmt.Fprintf(s.ErrStream, "現在設定されているPATHは%sです\n", Fpath)
	}

	if effective {
		if Fpath == "" {
			fmt.Fprintln(s.ErrStream, "対象ファイルが設定されていないので、プロキシを有効化できません")
			return ExitCodeExeFlagError
		}
		err := fproxy.SwitchProxyAuto(Fpath)
		if err != nil {
			fmt.Fprintf(s.ErrStream, "Failed to Comment Out.　自動コメントアウトに失敗しました。エラーの原因:%s\n", err)
			return ExitCodeExeFlagError
		}
		fmt.Fprintln(s.ErrStream, "対象ファイルのコメントをはずし、プロキシを有効化しました\n")

	}
	if ineffective {
		if Fpath == "" {
			fmt.Fprintln(s.ErrStream, "対象ファイルが設定されていないので、プロキシを無効化できません")
			return ExitCodeExeFlagError
		}
		err := fproxy.SwitchProxyAuto(Fpath)
		if err != nil {
			fmt.Fprintf(s.ErrStream, "Failed to Comment Out.　自動コメントアウトに失敗しました。エラーの原因:%s\n", err)
			return ExitCodeExeFlagError
		}
		fmt.Fprintln(s.ErrStream, "対象ファイルにコメントをつけ、プロキシを無効化しました")
	}

	return ExitCodeOK
}

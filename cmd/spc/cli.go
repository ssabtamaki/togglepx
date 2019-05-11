package cli

import (
	"flag"
	"fmt"
	"github.com/ssabcire/switchpx/fproxy"
	"io"
	"net"
)

const (
	exitCodeOK = iota
	exitCodeParseFlagError
	exitCodeExeFlagError
	switchProxyError
	writeJsonFileError
)

type Stream struct {
	OutStream, ErrStream io.Writer
}

func (s *Stream) Run(args []string, p *PathIPConfig) int {
	// switch proxy cli
	flags := flag.NewFlagSet("spc", flag.ContinueOnError)
	flags.SetOutput(s.ErrStream)

	//変数名あとで変更しておく
	var orgP string //Valueの値をうまく設定しないといけない。
	flags.StringVar(&orgP, "pxip", "", "test")
	var checkIP bool
	flags.BoolVar(&checkIP, "checkip", false, "登録したネットワークアドレスの値を確認します")
	var cancelIP bool
	flags.BoolVar(&cancelIP, "cancelip", false, "登録したネットワークアドレスの値を解除します")
	//変数名あとで変更しておく
	var orgF string
	flags.StringVar(&orgF, "fpath", "", "test")
	var checkPath bool
	flags.BoolVar(&checkPath, "checkpath", false, "test")
	var cancelPath bool
	flags.BoolVar(&cancelPath, "cancelpath", false, "test")
	var switching bool
	flags.BoolVar(&switching, "switch", false, "test")

	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeParseFlagError
	}
	if orgP != "" {
		if net.ParseIP(orgP) != nil {
			p.PxIP = orgP
			fmt.Fprintf(s.ErrStream, "ネットワークアドレスを<%s>に設定しました\n", p.PxIP)
		} else {
			fmt.Fprint(s.ErrStream, "入力されたアドレスが正しくありません\n")
			return exitCodeExeFlagError
		}
	}
	if checkIP {
		if p.PxIP != "" {
			fmt.Fprintf(s.ErrStream, "config.jsonに設定されているネットワークアドレスは<%s>です\n", p.PxIP)
		} else {
			fmt.Fprint(s.ErrStream, "config.jsonにネットワークアドレスが登録されていません\n")
		}
	}
	if cancelIP {
		p.PxIP = ""
		fmt.Fprint(s.ErrStream, "設定されていたネットワークアドレスを解除しました\n")
	}
	if orgF != "" {
		p.FilePath = orgF
		fmt.Fprint(s.ErrStream, "自動で切り替える対象のファイルのパスを<%s>にしました\n", p.FilePath)
	}
	if checkPath {
		if p.FilePath != "" {
			fmt.Fprintf(s.ErrStream, "config.jsonに設定されているファイルのパスは<%s>です\n", p.FilePath)
		} else {
			fmt.Fprint(s.ErrStream, "config.jsonにファイルのパスが登録されていません\n")
		}
	}
	if cancelPath {
		p.FilePath = ""
		fmt.Fprint(s.ErrStream, "設定されていたパスを解除しました\n")
	}
	if switching {
		err := fproxy.SwitchProxyAuto(p.FilePath)
		if err != nil {
			fmt.Fprint(s.ErrStream, "プロキシの切り替えに失敗しました。設定ファイルのパスが正しいか確認してください。")
			return switchProxyError
		}
		fmt.Fprint(s.ErrStream, "プロキシを切り替えました\n")
	}

	err := writeToJsonFile(p)
	if err != nil {
		return writeJsonFileError
	}

	return exitCodeOK
}

package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"togglepx/lib"
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

func (s *Stream) Run(args []string, p *lib.PathIPConfig) int {
	// taggle proxy auto
	flags := flag.NewFlagSet("togglepx", flag.ContinueOnError)
	flags.SetOutput(s.ErrStream)

	var tempP string
	flags.StringVar(&tempP, "pxip", "", "ネットワークアドレスを設定します。プロキシ下のネットワークアドレスをここに設定してください。")
	var checkIP bool
	flags.BoolVar(&checkIP, "checkip", false, "設定したネットワークアドレスの値を確認します")
	var cancelIP bool
	flags.BoolVar(&cancelIP, "cancelip", false, "設定したネットワークアドレスの値を解除します")
	var tempF string
	flags.StringVar(&tempF, "filepath", "", "PATHを設定できます。プロキシが書かれたdotfileのPATHを記載してください")
	var checkPath bool
	flags.BoolVar(&checkPath, "checkpath", false, "設定されたPATHの値を確認します")
	var cancelPath bool
	flags.BoolVar(&cancelPath, "cancelpath", false, "設定されたPATHの値を解除します")
	var switching bool
	flags.BoolVar(&switching, "switch", false, "プロキシをコメント/コメントアウトして、切り替えます")

	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeParseFlagError
	}
	if tempP != "" {
		if net.ParseIP(tempP) != nil {
			p.PxIP = tempP
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
	if tempF != "" {
		p.FilePath = tempF
		fmt.Fprintf(s.ErrStream, "自動で切り替える対象のファイルのパスを<%s>にしました\n", p.FilePath)
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
		err := lib.ToggleProxyAuto(p.FilePath)
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

func writeToJsonFile(p *lib.PathIPConfig) error {
	json, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(lib.JsonPath, json, 0666)
	if err != nil {
		return err
	}
	return nil
}
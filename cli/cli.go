package cli

import (
	"flag"
	"fmt"
	"io"
	"net"
)

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
	ExitCodeExeFlagError
	JsonError
)

type Stream struct {
	OutStream, ErrStream io.Writer
}

func (s *Stream) Run(args []string, p *PathIPConfig) int {
	err := p.ReadJsonTransfer()
	if err != nil {
		fmt.Print("Jsonファイルから構造体への変換に失敗しました。")
		return JsonError
	}


	flags := flag.NewFlagSet("sfp", flag.ContinueOnError)
	flags.SetOutput(s.ErrStream)

	var orgP string
	flags.StringVar(&orgP, "pxip", p.PxIP, "test")
	if net.ParseIP(orgP) == nil {
		//エラーでExit
	} else {
		p.PxIP = orgP
		/*Jsonに書き込む処理*/
	}

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	return ExitCodeOK
}



func JsonTest(p *PathIPConfig) (*PathIPConfig, error) {
	err := p.ReadJsonTransfer()
	if err != nil {
		fmt.Print("Jsonファイルから構造体への変換に失敗しました。")
		return nil, err
	}
	return p, nil
}

/*
//CLIのコマンドの設定と実行
func (s *Stream) Run(args []string) int {
	//sfp:switch file proxy
	flags := flag.NewFlagSet("sfp", flag.ContinueOnError)
	flags.SetOutput(s.ErrStream)

	pathIPConfig := PathIPConfig{}
	//flags.Var(&pathIPConfig.PxIP, "pxip", "ネットワークアドレスを登録します。登録するときは、プロキシ環境下のアドレスを登録してください")
	flags.StringVar(&pathIPConfig.PxIP, "pxip", pathIPConfig.PxIP, "ネットワークアドレスを登録します。登録するときは、プロキシ環境下のアドレスを登録してください")

	var checkIP bool
	flags.BoolVar(&checkIP, "checkip", false, "登録したネットワークアドレスの値を確認します")
	var cancelIP bool
	flags.BoolVar(&cancelIP, "cancelip", false, "登録したネットワークアドレスの値を解除します")

	flags.StringVar(&pathIPConfig.FilePath, "filepath", pathIPConfig.FilePath, "プロキシのオンオフ対象のファイルのPATHを設定します")
	var checkPath bool
	flags.BoolVar(&checkPath, "checkpath", false, "現在設定されているパスを確認します")
	var cancelPath bool
	flags.BoolVar(&cancelPath, "cancelpath", false, "登録されているパスを解除します")

	var switching bool
	flags.BoolVar(&switching, "switch", false, "対象ファイルにコメントされているときは、対象ファイルのコメントをはずし、プロキシを有効化します。"+
		"対象ファイルにコメントがされていないときは、対象ファイルにコメントをつけ、プロキシを無効化します")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if len(pathIPConfig.PxIP) > 1 {
		parse := net.ParseIP(pathIPConfig.PxIP)
		if parse == nil {
			fmt.Fprint(s.ErrStream, "入力されたIPアドレスが正しくありません")
			return ExitCodeExeFlagError
		}
		pxIP := parse.String()
		pathIPConfig.PxIP = pxIP		//&pathIPConfig.PxIP = &pxIP
		// 以下で設定ファイルに書き込む
	}

	if checkIP {
		//設定ファイルを読み込む
		if PxIP == "" {
			fmt.Fprint(s.ErrStream, "現在設定されているネットワークアドレスはありません")
			return ExitCodeExeFlagError
		}
		fmt.Printf("現在設定されているネットワークアドレスは%sです", )
	}
	if cancelIP {
		//設定ファイルを読み込む
		PxIP = ""
		//設定ファイルに書き込む
		fmt.Print("設定されたネットワークアドレスを解除しました")
	}

	if Fpath != "" {
		fmt.Fprintf(s.ErrStream, "プロキシのオンオフ対象のファイルを、%sにPATHを設定しました\n", Fpath)
	}
	if checkPath {
		fmt.Fprintf(s.ErrStream, "現在設定されているPATHは%sです\n", Fpath)
	}
	if cancelPath {
		Fpath = ""
		fmt.Fprintln(s.ErrStream, "設定されているPATHを取り消しました")
	}

	if switching {
		if Fpath == "" {
			fmt.Fprintln(s.ErrStream, "対象ファイルが設定されていないので、プロキシを有効化できません")
			return ExitCodeExeFlagError
		}
		err := fproxy.SwitchProxyAuto(Fpath)
		if err != nil {
			fmt.Fprintf(s.ErrStream, "自動コメントアウトに失敗しました。エラーの原因:%s\n", err)
			return 3 //ExitCodeExeFlagError
		}
		fmt.Fprintln(s.ErrStream, "対象ファイルのコメントをはずし、プロキシを有効化しました")
	}

	return ExitCodeOK
}
*/


//type IP net.IP
//flag.Valueインタフェース実装
//func (p *PathIPConfig) String() string {
//	return p.String()
//}
//func (p *PathIPConfig) Set(s string) error {
//	parse := net.ParseIP(s)
//	if parse == nil {
//		return errors.New("入力されたIPアドレスが正しくありません")
//	}
//	p.PxIP = parse.String()		//ここポインタいらない？
//
//	/* 設定ファイルに書き込む */
//
//	fmt.Fprintf(os.Stderr, "ネットワークアドレス%sをセットしました\n", i.String())
//	return nil
//}
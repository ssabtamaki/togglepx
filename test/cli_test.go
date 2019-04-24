package test

import (
	"stepupgo/cli"
	"testing"
)

func Test_jsonInclude(t *testing.T) {
	pathIPConfig := &cli.PathIPConfig{}
	p, err := cli.JsonTest(pathIPConfig)
	if err != nil {
		t.Errorf("JsonTestの実行に失敗")
	}
	expected := "test"
	if expected != p.FilePath {
		t.Errorf("want:%v, got:%v", expected, p.FilePath)
	}
	expected = "127.0.0.1"
	if expected != p.PxIP {
		t.Errorf("want:%v, got:%v", expected, p.PxIP)
	}
}


//func Test_cli(t *testing.T) {
//	cli.PxIP = cli.IP(ip)
//	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
//	stream := &cli.Stream{OutStream: outStream, ErrStream: errStream}
//	args := strings.Split("sfp -checkip", " ")
//	status := stream.Run(args)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected := fmt.Sprintf("現在設定されているネットワークアドレスは%sです\n", cli.PxIP.String())
//	if !strings.Contains(errStream.String(), expected) {
//		t.Errorf("output=%q, want %q", errStream.String(), expected)
//	}
//
//	errStream.Reset()
//	args = strings.Split("sfp -cancelip", " ")
//	status = stream.Run(args)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected = fmt.Sprintln("設定されているネットワークアドレスを取り消しました")
//	if !strings.Contains(errStream.String(), expected) {
//		t.Errorf("output=%q, want %q", errStream.String(), expected)
//	}
//
//	errStream.Reset()
//	args = strings.Split("sfp -filepath proxy.txt", " ")
//	status = stream.Run(args)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected = fmt.Sprintf("プロキシのオンオフ対象のファイルを、%sにPATHを設定しました\n", filePath)
//	if !strings.Contains(errStream.String(), expected) {
//		t.Errorf("output=%q, want %q", errStream.String(), expected)
//	}
//
//	errStream.Reset()
//	args = strings.Split("sfp -checkpath", " ")
//	status = stream.Run(args)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected = fmt.Sprintf("現在設定されているPATHは%sです\n", filePath)
//	if !strings.Contains(errStream.String(), expected) {
//		t.Errorf("output=%q, want %q", errStream.String(), expected)
//	}
//
//	errStream.Reset()
//	args = strings.Split("sfp -cancelpath", " ")
//	status = stream.Run(args)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	if cli.Fpath != "" {
//		t.Errorf("Fpath error")
//	}
//	expected = fmt.Sprintln("設定されているPATHを取り消しました")
//	if !strings.Contains(errStream.String(), expected) {
//		t.Errorf("output=%q, want %q", errStream.String(), expected)
//	}
//
//	errStream.Reset()
//	cli.Fpath = filePath
//	args = strings.Split("sfp -switch", " ")
//	status = stream.Run(args)
//	cli.PxIP = cli.IP(net.IP(ip))
//	//以下で、ファイルがきちんと切り替わっているかチェック
//
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//
//	//pxip以外、すべてテスト完了
//
//}

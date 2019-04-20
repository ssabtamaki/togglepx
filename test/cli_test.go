package test

import (
	"bytes"
	"fmt"
	"net"
	"stepupgo/cli"
	"strings"
	"testing"
)

const (
	filePath = "/test/cli"
	ip       = "127.0.0.1"
)

func Test_swfpx(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	stream := &cli.Stream{OutStream: outStream, ErrStream: errStream}
	/*
		args := strings.Split("swfpx -pxip 127.0.0.1", " ")
		status := stream.Run(args)
		if status != cli.ExitCodeOK {
			t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
		}
		expected := fmt.Sprintf("ネットワークアドレス%sを登録しました\n", ip)
		if !strings.Contains(errStream.String(), expected) {
			t.Errorf("output=%q, want %q", errStream.String(), expected)
		}
	*/
	errStream.Reset()
	args := strings.Split("swfpx -checkip", " ")
	status := stream.Run(args)
	if status != cli.ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
	}
	expected := fmt.Sprintf("現在設定されているネットワークアドレスは%sです\n", cli.PxIP)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -cancelip", " ")
	status = stream.Run(args)
	if status != cli.ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
	}
	expected = fmt.Sprintln("設定されているネットワークアドレスを取り消しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -filepath /test/cli", " ")
	status = stream.Run(args)
	if status != cli.ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
	}
	expected = fmt.Sprintf("プロキシのオンオフ対象のファイルを、%sにPATHを設定しました\n", filePath)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -cancelpath", " ")
	status = stream.Run(args)
	if status != cli.ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
	}
	expected = fmt.Sprintln("設定されているPATHを取り消しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -checkpath", " ")
	status = stream.Run(args)
	if status != cli.ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
	}
	expected = fmt.Sprintf("現在設定されているPATHは%sです\n", filePath)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -switch", " ")
	status = stream.Run(args)
	cli.Fpath = "/Users/ssab/go/src/stepupgo/test/proxy.txt"
	cli.PxIP = cli.IP(net.IP(ip))
	if status != cli.ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
	}


}

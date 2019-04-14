package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const filePath = "/test/cli"
const ip = "127.0.0.1"

func Test_swfpx(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	stream := &Stream{OutStream: outStream, ErrStream: errStream}
	args := strings.Split("swfpx -pxip 127.0.0.1", " ")
	status := stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected := fmt.Sprintf("ネットワークアドレス%sを登録しました", ip)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -checkip", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("現在設定されているネットワークアドレスは%sです", PxIP)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -cancelip", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("設定されているネットワークアドレスを取り消しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -filepath /test/cli", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("プロキシのオンオフ対象のファイルを、%sにPATHを設定しました", filePath)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -cancelpath", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("設定されているPATHを取り消しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -checkpath", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("現在設定されているPATHは%sです", filePath)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -effective", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("対象ファイルのコメントをはずし、プロキシを有効化します")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -ineffective", " ")
	status = stream.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("対象ファイルにコメントをつけ、プロキシを無効化しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}
}
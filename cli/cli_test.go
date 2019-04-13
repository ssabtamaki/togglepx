package cli

import (
	"bytes"
	"strings"
	"testing"
	"fmt"
)

const FilePath = "/test/cli"

func Test_swfpx(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &cli{outStream:outStream, errStream: errStream}
	args := strings.Split("swfpx -fpath /test/cli", " ")
	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected := fmt.Sprintf("プロキシのオンオフ対象のファイルを、%sにPATHを設定しました", FilePath)
	//ここちょっと問題がある。fpathをグローバルにするとテストが通るが、ローカルだと通らない
	//原因を救命する必要がある
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -cancel", " ")
	status = cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("設定されているPATHを取り消しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -check", " ")
	status = cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("現在設定されているPATHは%sです", FilePath)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -effective", " ")
	status = cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("対象ファイルのコメントをはずし、プロキシを有効化します")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}

	errStream.Reset()
	args = strings.Split("swfpx -ineffective", " ")
	status = cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected = fmt.Sprintf("対象ファイルにコメントをつけ、プロキシを無効化しました")
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}
}

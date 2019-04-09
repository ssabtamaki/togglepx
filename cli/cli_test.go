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
	args := strings.Split("swfpx -filepath /test/cli", " ")
	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus = %d, want %d", status, ExitCodeOK)
	}
	expected := fmt.Sprintf(FilePath)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("output=%q, want %q", errStream.String(), expected)
	}
}

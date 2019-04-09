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

func (c *cli) Run(args []string) int {
	//switch file proxy = swfpx
	flags := flag.NewFlagSet("swfpx", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	var filepath string
	flags.StringVar(&filepath, "filepath", filepath, "specify filepath for proxy")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if filepath != "" {
		fmt.Fprint(c.errStream, filepath)
	}
	return ExitCodeOK
}
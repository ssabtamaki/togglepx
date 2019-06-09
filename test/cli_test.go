//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"tpa/lib"
//	"strings"
//	"testing"
//)
//
//func Test_CliSubCommand_pxip(t *testing.T) {
//	p := cli.PathIPConfig{}
//	err := p.ReadJsonTransfer(cli.JsonPath)
//	if err != nil {
//		t.Errorf("transfer err")
//	}
//
//	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
//	stream := &cli.Stream{OutStream: outStream, ErrStream: errStream}
//	args := strings.Split("sfp -pxip 192.168.11.0", " ")
//	status := stream.Run(args, &p)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected := "192.168.11.0"
//	if expected != p.PxIP {
//		t.Errorf("want %v, get %v", expected, p.PxIP)
//	}
//	args = strings.Split("sfp -pxip 192.168", " ")
//	status = stream.Run(args, &p)
//	if status != cli.ExitCodeExeFlagError {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	args = strings.Split("sfp -pxip", " ")
//	status = stream.Run(args, &p)
//	if status != cli.ExitCodeParseFlagError {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//}
//
//func Test_CliSubCommand_AllBool(t *testing.T) {
//	p := cli.PathIPConfig{}
//	err := p.ReadJsonTransfer(cli.JsonPath)
//	if err != nil {
//		t.Errorf("transfer err")
//	}
//	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
//	stream := &cli.Stream{OutStream: outStream, ErrStream: errStream}
//	args := strings.Split("sfp -checkip", " ")
//	status := stream.Run(args, &p)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected := fmt.Sprintf("config.jsonに設定されているネットワークアドレスは%sです", p.PxIP)
//	if !strings.Contains(errStream.String(), expected) {
//		t.Errorf("want %v, get %v", expected, errStream.String())
//	}
//
//	errStream.Reset()
//	args = strings.Split("sfp -cancelip", " ")
//	status = stream.Run(args, &p)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected = ""
//	if expected != p.PxIP {
//		t.Errorf("want %v, get %v", expected, p.PxIP)
//	}
//	p.PxIP = "127.0.0.1"
//
//	errStream.Reset()
//	args = strings.Split("sfp -cancelpath", " ")
//	status = stream.Run(args, &p)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	expected = ""
//	if expected != p.FilePath {
//		t.Errorf("want %v, get %v", expected, p.FilePath)
//	}
//
//	p.FilePath = "/Users/ssab/go/src/tpa/test/proxy.txt"
//	errStream.Reset()
//	args = strings.Split("sfp -switch", " ")
//	status = stream.Run(args, &p)
//	if status != cli.ExitCodeOK {
//		t.Errorf("ExitStatus = %d, want %d", status, cli.ExitCodeOK)
//	}
//	input, err := ioutil.ReadFile(p.FilePath)
//	if err != nil {
//		t.Error("Error to Read File")
//	}
//	expected = "# proxy=<>:<><>"
//	if !strings.Contains(string(input), expected) {
//		t.Errorf("get [%v], want [%v]", string(input), expected)
//	}
//}

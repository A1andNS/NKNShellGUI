package command

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// ExecCommand 命令执行分类
func ExecCommand(cmd string) string {
	switch runtime.GOOS {
	case "linux":
		result := linuxExec(cmd)
		return result
	case "windows":
		result := WinodwsExec(cmd)
		return result
	default:
		result := linuxExec(cmd)
		return result
	}
}

// WinodwsExec windows命令执行
func WinodwsExec(args ...string) string {
	args = append([]string{"/c"}, args...)
	cmd := exec.Command("cmd", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	outpip, err := cmd.StdoutPipe()
	defer outpip.Close()

	if err != nil {
		return "error"
	}
	err = cmd.Start()
	if err != nil {
		return "error"
	}
	out, err := ioutil.ReadAll(outpip)
	if err != nil {
		return "error"
	}
	cmdout := ConverByte2String(out, "GB18030")
	return string(cmdout)
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// ConverByte2String 转字节为字符串，在提取命令运行结果时，可以吧管道中的byte结果转为string输出
func ConverByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// linuxExec Linux命令执行
func linuxExec(args ...string) string {
	args = append([]string{"-c"}, args...)
	cmd := exec.Command(os.Getenv("SHELL"), args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	outpip, err := cmd.StdoutPipe()
	defer outpip.Close()
	if err != nil {
		return "error"
	}
	err = cmd.Start()
	if err != nil {
		return "error"
	}
	out, err := ioutil.ReadAll(outpip)
	if err != nil {
		return "error"
	}
	return string(out)
}

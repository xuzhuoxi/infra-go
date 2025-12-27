package cmdx

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	CmdExit    = "exit"
	CmdVersion = "version"
	Version    = "1.0.0"
)

type ICommandLineListener interface {
	// StartListen 开始监听
	StartListen()
	// StopListen 停止监听
	StopListen()

	// SetFrontTips 设置提示前缀
	SetFrontTips(tips string)
	// SetRepeatCount 设置重复次数
	SetRepeatCount(repeatCount int)
	// SetMatchCase  设置大小写敏感
	SetMatchCase(matchKeyCase bool, matchValueCase bool)
	// SetTrimInput  设置首尾空字符处理
	SetTrimInput(trimInput bool)

	// MapCommand 映射处理函数
	MapCommand(cmd string, f func(flagSet *FlagSetExtend, args []string))
}

func CreateCommandLineListener(frontTips string, repeatCount int) ICommandLineListener {
	rs := CommandLineListener{
		FrontTips:    frontTips,
		RepeatCount:  repeatCount,
		CurrentCount: 0,
		TrimInput:    true,

		handler:  make(map[string]func(flagSet *FlagSetExtend, args []string)),
		reader:   nil,
		exitFlag: false}
	rs.MapCommand(CmdExit, rs.confirmExit)
	rs.MapCommand(CmdVersion, version)
	return &rs
}

type CommandLineListener struct {
	FrontTips    string
	RepeatCount  int
	CurrentCount int

	MatchKeyCase   bool
	MatchValueCase bool
	TrimInput      bool

	handler  map[string]func(flagSet *FlagSetExtend, args []string)
	reader   *bufio.Reader
	exitFlag bool
}

func (c *CommandLineListener) StartListen() {
	c.reader = bufio.NewReader(os.Stdin)
	c.exitFlag = false
	c.nextCommand()
}

func (c *CommandLineListener) StopListen() {
	c.exitFlag = true
}

func (c *CommandLineListener) SetFrontTips(tips string) {
	c.FrontTips = tips
}

func (c *CommandLineListener) SetRepeatCount(repeatCount int) {
	c.RepeatCount = repeatCount
}

func (c *CommandLineListener) SetMatchCase(matchKeyCase bool, matchValueCase bool) {
	c.MatchKeyCase, c.MatchValueCase = matchKeyCase, matchValueCase
}

func (c *CommandLineListener) SetTrimInput(trimInput bool) {
	c.TrimInput = trimInput
}

func (c *CommandLineListener) MapCommand(cmd string, f func(flagSet *FlagSetExtend, args []string)) {
	c.handler[cmd] = f
}

//private----------

func (c *CommandLineListener) nextCommand() {
	if c.exitFlag {
		return
	}
	c.prepareCommand()
	c.listenCommand()
	exit := c.finishCommand()
	if !exit {
		time.Sleep(time.Millisecond * 200)
		c.nextCommand()
	}
}

func (c *CommandLineListener) prepareCommand() {
	fmt.Print(c.FrontTips)
}

func (c *CommandLineListener) listenCommand() {
	input, _ := c.reader.ReadString('\n') //定义一行输入的内容分隔符。
	if len(input) == 1 {
		return
	}
	if c.TrimInput {
		input = strings.TrimSpace(input)
	}
	cmdArgs := strings.Split(input, " ")
	cmd := strings.ToLower(cmdArgs[0])

	flagSet := NewFlagSetExtend(cmd, flag.ContinueOnError)
	f := c.handler[cmd]
	if nil == f {
		return
	}
	if !c.MatchKeyCase {
		for index := 1; index < len(cmdArgs); index += 2 {
			cmdArgs[index] = strings.ToLower(cmdArgs[index])
		}
	}
	if !c.MatchValueCase {
		for index := 2; index < len(cmdArgs); index += 2 {
			cmdArgs[index] = strings.ToLower(cmdArgs[index])
		}
	}
	f(flagSet, cmdArgs[1:])
}

func (c *CommandLineListener) finishCommand() bool {
	c.CurrentCount++
	unLimit := c.RepeatCount <= 0
	if unLimit || c.CurrentCount < c.RepeatCount {
		return false
	}
	return true
}

func (c *CommandLineListener) confirmExit(_ *FlagSetExtend, _ []string) {
	fmt.Print("Ary you sure to exit:")
	var input string
	fmt.Scanln(&input)
	inputTrim := strings.ToLower(strings.TrimSpace(input))
	if "yes" == inputTrim || "y" == inputTrim || "1" == inputTrim {
		c.exitFlag = true
	}
}

func version(_ *FlagSetExtend, _ []string) {
	fmt.Println("version=" + Version)
}

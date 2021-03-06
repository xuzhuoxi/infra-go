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
	StartListen()
	StopListen()
	SetFrontTips(tips string)
	SetRepeatCount(repeatCount int)
	MapCommand(cmd string, f func(flagSet *FlagSetExtend, args []string))
}

func CreateCommandLineListener(frontTips string, repeatCount int) ICommandLineListener {
	rs := CommandLineListener{frontTips, repeatCount, 0, make(map[string]func(flagSet *FlagSetExtend, args []string)), nil, false}
	rs.MapCommand(CmdExit, rs.confirmExit)
	rs.MapCommand(CmdVersion, version)
	return &rs
}

type CommandLineListener struct {
	FrontTips    string
	RepeatCount  int
	CurrentCount int
	handler      map[string]func(flagSet *FlagSetExtend, args []string)
	reader       *bufio.Reader
	exitFlag     bool
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
	inputTrim := strings.ToLower(strings.TrimSpace(input))
	cmdArgs := strings.Split(inputTrim, " ")
	cmd := cmdArgs[0]

	flagSet := NewFlagSetExtend(cmd, flag.ContinueOnError)
	f := c.handler[cmd]
	if nil == f {
		return
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

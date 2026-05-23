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
	// CmdExit is the built-in exit command name.
	// CmdExit 为内置退出命令名。
	CmdExit = "exit"
	// CmdVersion is the built-in version command name.
	// CmdVersion 为内置版本查询命令名。
	CmdVersion = "version"
	// Version is the package version string.
	// Version 为包版本号字符串。
	Version = "1.0.0"
)

// ICommandLineListener defines the contract for an interactive command-line listener.
// ICommandLineListener 定义交互式命令行监听器的契约。
type ICommandLineListener interface {
	// StartListen starts reading commands from stdin in a loop.
	// StartListen 开始循环从标准输入读取命令。
	StartListen()
	// StopListen stops the command listening loop.
	// StopListen 停止命令监听循环。
	StopListen()

	// SetFrontTips sets the prompt prefix printed before each command input.
	// SetFrontTips 设置每次输入命令前打印的提示前缀。
	SetFrontTips(tips string)
	// SetRepeatCount sets the maximum number of commands to process; non-positive means unlimited.
	// SetRepeatCount 设置最多处理的命令次数；非正数表示不限制。
	SetRepeatCount(repeatCount int)
	// SetMatchCase sets whether command keys and values are matched case-sensitively.
	// SetMatchCase 设置命令键与值是否区分大小写匹配。
	SetMatchCase(matchKeyCase bool, matchValueCase bool)
	// SetTrimInput sets whether leading and trailing whitespace is trimmed from input.
	// SetTrimInput 设置是否去除输入内容的首尾空白字符。
	SetTrimInput(trimInput bool)

	// MapCommand registers a handler for the given command name.
	// MapCommand 为指定命令名注册处理函数。
	MapCommand(cmd string, f func(flagSet *FlagSetExtend, args []string))
}

// CreateCommandLineListener creates a command-line listener with the given prompt and repeat count.
// CreateCommandLineListener 使用给定提示前缀与重复次数创建命令行监听器。
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

// CommandLineListener reads stdin commands and dispatches them to registered handlers.
// CommandLineListener 从标准输入读取命令并分派到已注册的处理函数。
type CommandLineListener struct {
	// FrontTips is the prompt printed before each command input.
	// FrontTips 为每次输入命令前打印的提示前缀。
	FrontTips string
	// RepeatCount is the maximum number of commands to process; non-positive means unlimited.
	// RepeatCount 为最多处理的命令次数；非正数表示不限制。
	RepeatCount int
	// CurrentCount is the number of commands processed so far.
	// CurrentCount 为当前已处理的命令次数。
	CurrentCount int

	// MatchKeyCase controls case-sensitive matching for command argument keys.
	// MatchKeyCase 控制命令参数键是否区分大小写匹配。
	MatchKeyCase bool
	// MatchValueCase controls case-sensitive matching for command argument values.
	// MatchValueCase 控制命令参数值是否区分大小写匹配。
	MatchValueCase bool
	// TrimInput controls whether leading and trailing whitespace is trimmed from input.
	// TrimInput 控制是否去除输入内容的首尾空白字符。
	TrimInput bool

	handler  map[string]func(flagSet *FlagSetExtend, args []string)
	reader   *bufio.Reader
	exitFlag bool
}

// StartListen starts reading commands from stdin in a loop.
// StartListen 开始循环从标准输入读取命令。
func (c *CommandLineListener) StartListen() {
	c.reader = bufio.NewReader(os.Stdin)
	c.exitFlag = false
	c.nextCommand()
}

// StopListen stops the command listening loop.
// StopListen 停止命令监听循环。
func (c *CommandLineListener) StopListen() {
	c.exitFlag = true
}

// SetFrontTips sets the prompt prefix printed before each command input.
// SetFrontTips 设置每次输入命令前打印的提示前缀。
func (c *CommandLineListener) SetFrontTips(tips string) {
	c.FrontTips = tips
}

// SetRepeatCount sets the maximum number of commands to process; non-positive means unlimited.
// SetRepeatCount 设置最多处理的命令次数；非正数表示不限制。
func (c *CommandLineListener) SetRepeatCount(repeatCount int) {
	c.RepeatCount = repeatCount
}

// SetMatchCase sets whether command keys and values are matched case-sensitively.
// SetMatchCase 设置命令键与值是否区分大小写匹配。
func (c *CommandLineListener) SetMatchCase(matchKeyCase bool, matchValueCase bool) {
	c.MatchKeyCase, c.MatchValueCase = matchKeyCase, matchValueCase
}

// SetTrimInput sets whether leading and trailing whitespace is trimmed from input.
// SetTrimInput 设置是否去除输入内容的首尾空白字符。
func (c *CommandLineListener) SetTrimInput(trimInput bool) {
	c.TrimInput = trimInput
}

// MapCommand registers a handler for the given command name.
// MapCommand 为指定命令名注册处理函数。
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

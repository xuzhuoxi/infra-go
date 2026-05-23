// Package cmdx provides interactive command-line listening and extended flag/argument parsing.
// Package cmdx 提供交互式命令行监听及扩展的标志与参数解析能力。
//
// Created by xuzhuoxi on 2019-06-07.
// @author xuzhuoxi
//
// Capabilities:
//   - CommandLineListener: read stdin commands in a loop, map command names to handlers with optional repeat limit
//   - FlagSetExtend: embeds flag.FlagSet and merges standard flags with key=value style arguments via ArgSet
//   - ArgSet: parse -key=value and --key=value arguments and retrieve typed values
//
// 主要职责：
//   - CommandLineListener：循环读取标准输入命令，将命令名映射到处理函数，支持重复次数限制
//   - FlagSetExtend：嵌入 flag.FlagSet，并通过 ArgSet 合并标准 flag 与 key=value 风格参数
//   - ArgSet：解析 -key=value 与 --key=value 参数并按类型取值
package cmdx

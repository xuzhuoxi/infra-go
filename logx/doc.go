// Package logx provides multi-sink logging with levels, file rotation, and named logger management.
// Package logx 提供多输出目标、分级日志、文件滚动切割及命名 Logger 管理能力。
//
// Created by xuzhuoxi on 2023-06-30.
// @author xuzhuoxi
//
// Capabilities:
//   - ILogger with console, rolling file, daily file, and daily-plus-rolling file outputs
//   - Log levels from Trace through Fatal; package-level helpers on a default global logger
//   - LogConfig and CfgLog for path, size, and level configuration; optional ANSI color prefixes
//   - ILoggerManager / LoggerManager: register, lookup, and dispatch to named loggers
//   - LoggerSupport embeddable helper for types that need injectable ILogger access
//
// 主要职责：
//   - ILogger 支持控制台、按体积分卷、按日期分卷及日期加体积分卷等输出方式
//   - Trace 至 Fatal 日志等级；包级函数委托给默认全局 Logger
//   - LogConfig 与 CfgLog 配置路径、大小与等级；可选 ANSI 颜色前缀
//   - ILoggerManager / LoggerManager：注册、查找命名 Logger 并按名分发日志
//   - LoggerSupport 可嵌入类型，便于注入与获取 ILogger
package logx

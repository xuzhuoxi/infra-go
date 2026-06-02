// Package regexpx provides precompiled regular-expression pattern strings for validation and matching.
// Package regexpx 提供预置的正则表达式模式字符串，用于校验与匹配。
//
// Patterns cover numeric types, common text (email, URL, Chinese ID), POSIX ASCII classes,
// Perl shorthands, and Unicode general/script classes. Combine with regexp.Compile or regexp.MatchString.
//
// 主要职责：
//   - regexp.go：数字、浮点、字母、邮箱、URL、身份证等常用校验模式
//   - ascii.go：POSIX 字符类（[[:alnum:]] 等）
//   - perl.go：Perl 风格元字符与转义（\d、\w、^、$ 等）
//   - unicode.go：Unicode 通用类别与各书写系统（\p{Han} 等）
//
// See https://en.wikipedia.org/wiki/Regular_expression for POSIX BRE/ERE background.
package regexpx

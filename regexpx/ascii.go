package regexpx

// ASCII类
const (
	AscAlnum  = `[[:alnum:]]`  // 字母数字 (相当于 [0-9A-Za-z])
	AscAlpha  = `[[:alpha:]]`  // 字母 (相当于 [A-Za-z])
	AscAscii  = `[[:ascii:]]`  // ASCII 字符集 (相当于 [\x00-\x7F])
	AscBlank  = `[[:blank:]]`  // 空白占位符 (相当于 [\t ])
	AscCntrl  = `[[:cntrl:]]`  // 控制字符 (相当于 [\x00-\x1F\x7F])
	AscDigit  = `[[:digit:]]`  // 数字 (相当于 [0-9])
	AscGraph  = `[[:graph:]]`  // 图形字符 (相当于 [!-~])
	AscLower  = `[[:lower:]]`  // 小写字母 (相当于 [a-z])
	AscPrint  = `[[:print:]]`  // 可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
	AscPunct  = `[[:punct:]]`  // 标点符号 (相当于 [!-/:-@[-反引号{-~])
	AscSpace  = `[[:space:]]`  // 空白字符(相当于 [\t\n\v\f\r ])
	AscUpper  = `[[:upper:]]`  // 大写字母(相当于 [A-Z])
	AscWord   = `[[:word:]]`   // 单词字符(相当于 [0-9A-Za-z_])
	AscXdigit = `[[:xdigit:]]` // 16進制字符集(相当于 [0-9A-Fa-f])
)

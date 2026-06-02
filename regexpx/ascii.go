package regexpx

// ASCII类
const (
	AscAlnum  = `[[:alnum:]]`  // AscAlnum matches POSIX letters and digits. 字母数字（相当于 [0-9A-Za-z]）
	AscAlpha  = `[[:alpha:]]`  // AscAlpha matches POSIX letters. 字母（相当于 [A-Za-z]）
	AscAscii  = `[[:ascii:]]`  // AscAscii matches the ASCII character set. ASCII 字符集（相当于 [\x00-\x7F]）
	AscBlank  = `[[:blank:]]`  // AscBlank matches horizontal blank characters. 空白占位符（相当于 [\t ]）
	AscCntrl  = `[[:cntrl:]]`  // AscCntrl matches control characters. 控制字符（相当于 [\x00-\x1F\x7F]）
	AscDigit  = `[[:digit:]]`  // AscDigit matches POSIX digits. 数字（相当于 [0-9]）
	AscGraph  = `[[:graph:]]`  // AscGraph matches visible (graphic) characters. 图形字符（相当于 [!-~]）
	AscLower  = `[[:lower:]]`  // AscLower matches lowercase letters. 小写字母（相当于 [a-z]）
	AscPrint  = `[[:print:]]`  // AscPrint matches printable characters. 可打印字符（相当于 [ -~]，含 [:graph:]）
	AscPunct  = `[[:punct:]]`  // AscPunct matches punctuation. 标点符号
	AscSpace  = `[[:space:]]`  // AscSpace matches whitespace. 空白字符（相当于 [\t\n\v\f\r ]）
	AscUpper  = `[[:upper:]]`  // AscUpper matches uppercase letters. 大写字母（相当于 [A-Z]）
	AscWord   = `[[:word:]]`   // AscWord matches word characters. 单词字符（相当于 [0-9A-Za-z_]）
	AscXdigit = `[[:xdigit:]]` // AscXdigit matches hexadecimal digits. 十六进制字符（相当于 [0-9A-Fa-f]）
)

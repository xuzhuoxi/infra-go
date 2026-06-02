package regexpx

// Perl类
const (
	PerlDigit    = `\d` // PerlDigit matches digits. 数字（相当于 [0-9]）
	PerlNotDigit = `\D` // PerlNotDigit matches non-digits. 非数字（相当于 [^0-9]）
	PerlSpace    = `\s` // PerlSpace matches whitespace. 空白（相当于 [\t\n\f\r ]）
	PerlNotSpace = `\S` // PerlNotSpace matches non-whitespace. 非空白（相当于 [^\t\n\f\r ]）
	PerlWord     = `\w` // PerlWord matches word characters. 单词字符（相当于 [0-9A-Za-z_]）
	PerlNotWord  = `\W` // PerlNotWord matches non-word characters. 非单词字符（相当于 [^0-9A-Za-z_]）

	PerlStart = `^`  // PerlStart matches start of line (with m) or start of text. 行首/文本开头（m 标记影响行为）
	PerlEnd   = `$`  // PerlEnd matches end of line (with m) or end of text. 行尾/文本结尾（m 标记影响行为）
	P_A       = `\A` // P_A matches start of entire text, ignoring m. 整个文本开头（忽略 m）
	P_b       = `\b` // P_b matches a word boundary. 单词边界
	P_B       = `\B` // P_B matches a non-word boundary. 非单词边界
	P_z       = `\z` // P_z matches end of entire text, ignoring m. 整个文本结尾（忽略 m）

	P_a         = `\a`         // P_a matches BEL (bell, \x07). 响铃符（相当于 \x07）
	P_backspace = `\x08`       // P_backspace matches backspace; use \x08 not \b. 退格符（须用 \x08，\b 为单词边界）
	P_f         = `\f`         // P_f matches form feed (\x0C). 换页符（相当于 \x0C）
	P_t         = `\t`         // P_t matches horizontal tab (\x09). 横向制表符（相当于 \x09）
	P_n         = `\n`         // P_n matches newline (\x0A). 换行符（相当于 \x0A）
	P_r         = `\r`         // P_r matches carriage return (\x0D). 回车符（相当于 \x0D）
	P_v         = `\v`         // P_v matches vertical tab (\x0B). 纵向制表符（相当于 \x0B）
	P_s8_3      = `\123`       // P_s8_3 matches an octal character (3 digits). 八进制编码字符（须 3 位数字）
	P_s16_3     = `\x7F`       // P_s16_3 matches a hex character (3 digits). 十六进制编码字符（须 3 位数字）
	P_s16       = `\x{10FFFF}` // P_s16 matches a hex Unicode code point (max 10FFFF). 十六进制 Unicode 码点（最大 10FFFF）
	P_Between   = `\Q...\E`    // P_Between matches literal text between \Q and \E. 匹配 \Q 与 \E 之间的字面文本

	//分组可以设置标记
	//i 不区分大小写 (默认为 false)
	//m 多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)
	//s 让 . 匹配 \n (默认为 false)
	//U 非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)
)

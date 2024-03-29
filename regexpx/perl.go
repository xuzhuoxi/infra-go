package regexpx

// Perl类
const (
	PerlDigit    = `\d` // 数字 (相当于 [0-9])
	PerlNotDigit = `\D` // 非数字 (相当于 [^0-9])
	PerlSpace    = `\s` // 空白 (相当于 [\t\n\f\r ])
	PerlNotSpace = `\S` // 非空白 (相当于[^\t\n\f\r ])
	PerlWord     = `\w` // 单词字符 (相当于 [0-9A-Za-z_])
	PerlNotWord  = `\W` // 非单词字符 (相当于 [^0-9A-Za-z_])

	PerlStart = `^`  // 如果标记 m=true 则匹配行首，否则匹配整个文本的开头（m 默认为 false）
	PerlEnd   = `$`  // 如果标记 m=true 则匹配行尾，否则匹配整个文本的结尾（m 默认为 false）
	P_A       = `\A` // 匹配整个文本的开头，忽略 m 标记
	P_b       = `\b` // 匹配单词边界
	P_B       = `\B` // 匹配非单词边界
	P_z       = `\z` // 匹配整个文本的结尾，忽略 m 标记

	P_a         = `\a`         //匹配响铃符（相当于 \x07）
	P_backspace = `\x08`       //注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，可以使用 \x08 表示退格符。
	P_f         = `\f`         //匹配换页符 （相当于 \x0C）
	P_t         = `\t`         //匹配横向制表符（相当于 \x09）
	P_n         = `\n`         //匹配换行符 （相当于 \x0A）
	P_r         = `\r`         //匹配回车符 （相当于 \x0D）
	P_v         = `\v`         //匹配纵向制表符（相当于 \x0B）
	P_s8_3      = `\123`       //匹配 8  進制编码所代表的字符（必须是 3 位数字）
	P_s16_3     = `\x7F`       //匹配 16 進制编码所代表的字符（必须是 3 位数字）
	P_s16       = `\x{10FFFF}` //匹配 16 進制编码所代表的字符（最大值 10FFFF）
	P_Between   = `\Q...\E`    //匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法

	//分组可以设置标记
	//i 不区分大小写 (默认为 false)
	//m 多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)
	//s 让 . 匹配 \n (默认为 false)
	//U 非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)
)

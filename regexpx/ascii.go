package regexpx

// ASCII类
const (
	ASC_alnum  = `[[:alnum:]]`  // 字母数字 (相当于 [0-9A-Za-z])
	ASC_alpha  = `[[:alpha:]]`  // 字母 (相当于 [A-Za-z])
	ASC_ascii  = `[[:ascii:]]`  // ASCII 字符集 (相当于 [\x00-\x7F])
	ASC_blank  = `[[:blank:]]`  // 空白占位符 (相当于 [\t ])
	ASC_cntrl  = `[[:cntrl:]]`  // 控制字符 (相当于 [\x00-\x1F\x7F])
	ASC_digit  = `[[:digit:]]`  // 数字 (相当于 [0-9])
	ASC_graph  = `[[:graph:]]`  // 图形字符 (相当于 [!-~])
	ASC_lower  = `[[:lower:]]`  // 小写字母 (相当于 [a-z])
	ASC_print  = `[[:print:]]`  // 可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
	ASC_punct  = `[[:punct:]]`  // 标点符号 (相当于 [!-/:-@[-反引号{-~])
	ASC_space  = `[[:space:]]`  // 空白字符(相当于 [\t\n\v\f\r ])
	ASC_upper  = `[[:upper:]]`  // 大写字母(相当于 [A-Z])
	ASC_word   = `[[:word:]]`   // 单词字符(相当于 [0-9A-Za-z_])
	ASC_xdigit = `[[:xdigit:]]` // 16進制字符集(相当于 [0-9A-Fa-f])
)

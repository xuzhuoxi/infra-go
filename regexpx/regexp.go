package regexpx

const (
	// Zero matches a single digit zero.
	// Zero 匹配单个数字 0。
	Zero = `^0$`
	// Digital matches a single decimal digit.
	// Digital 匹配单个十进制数字。
	Digital = `^[0-9]$`

	// PInt matches positive integers.
	// PInt 匹配正整数。
	PInt = `^[1-9]\d*$`
	// NInt matches negative integers.
	// NInt 匹配负整数。
	NInt = `^-[1-9]\d*$`
	// NNInt matches non-negative integers (positive integers and zero).
	// NNInt 匹配非负整数（正整数与 0）。
	NNInt = `^[1-9]\d*|0$`
	// NPInt matches non-positive integers (negative integers and zero).
	// NPInt 匹配非正整数（负整数与 0）。
	NPInt = `^-[1-9]\d*|0$`
	// Int matches integers.
	// Int 匹配整数。
	Int = `^-?[1-9]\d*|0$`

	// PFloat matches positive floating-point numbers.
	// PFloat 匹配正浮点数。
	PFloat = `^[1-9]\d*.\d*|0.\d*[1-9]\d*$`
	// NFloat matches negative floating-point numbers.
	// NFloat 匹配负浮点数。
	NFloat = `^-([1-9]\d*.\d*|0.\d*[1-9]\d*)$`
	// NNFloat matches non-negative floating-point numbers (positive floats and zero).
	// NNFloat 匹配非负浮点数（正浮点数与 0）。
	NNFloat = `^[1-9]\d*.\d*|0.\d*[1-9]\d*|0?.0+|0$`
	// NPFloat matches non-positive floating-point numbers (negative floats and zero).
	// NPFloat 匹配非正浮点数（负浮点数与 0）。
	NPFloat = `^(-([1-9]\d*.\d*|0.\d*[1-9]\d*))|0?.0+|0$`
	// Float matches floating-point numbers.
	// Float 匹配浮点数。
	Float = `^-?([1-9]\d*.d*|0.\d*[1-9]\d*|0?.0+|0)$`
)

const (
	// LowerLetter matches lowercase ASCII letters.
	// LowerLetter 匹配小写字母。
	LowerLetter = `[a-z]`
	// UpperLetter matches uppercase ASCII letters.
	// UpperLetter 匹配大写字母。
	UpperLetter = `[A-Z]`
	// Letter matches ASCII letters.
	// Letter 匹配字母。
	Letter = `[a-zA-Z]`
	// Space matches whitespace characters.
	// Space 匹配空白字符。
	Space = `[\t\n\f\r]`
	// Work matches word characters (digits, letters, underscore).
	// Work 匹配单词字符（数字、字母、下划线）。
	Work = `[0-9A-Za-z_]`

	// Chinese matches Han script characters.
	// Chinese 匹配中文字符（\p{Han}）。
	Chinese = `[\p{Han}]`
	// DoubleChar matches double-byte (non-ASCII) characters including CJK.
	// DoubleChar 匹配双字节字符（含汉字等非 ASCII）。
	DoubleChar = `[^x00-xff]`
	// EmptyLine matches blank lines.
	// EmptyLine 匹配空白行。
	EmptyLine = `ns*r`
	// HTML matches simple HTML tags (not reliable for deeply nested markup).
	// HTML 匹配简单 HTML 标记（复杂嵌套标记无法完全覆盖）。
	HTML = `<(S*?)[^>]*>.*?|<.*? /> `
	// EmptyHeadTail matches leading or trailing whitespace.
	// EmptyHeadTail 匹配首尾空白字符。
	EmptyHeadTail = `^\s*|\s*$ `
	// EMail matches email addresses.
	// EMail 匹配 Email 地址。
	EMail = `w[-w.+]*@([A-Za-z0-9][-A-Za-z0-9]+.)+[A-Za-z]{2, 14}`
	// EMail2 matches email addresses (alternate pattern).
	// EMail2 匹配 Email 地址（备用模式）。
	EMail2 = `w+([-+.]w+)*@w+([-.]w+)*.w+([-.]w+)*`
	// URL matches URLs with optional http/https/ftp/rtsp/mms scheme.
	// URL 匹配网址 URL（可选 http/https/ftp/rtsp/mms 协议）。
	URL = `^((https|http|ftp|rtsp|mms)?://)[^s]+`
	// URL2 matches URLs with a scheme prefix.
	// URL2 匹配带协议前缀的网址 URL。
	URL2 = `[a-zA-z]+://[^s]*`
	// Account matches usernames: letter first, 5–16 alphanumeric or underscore.
	// Account 匹配账号：字母开头，5–16 位字母、数字或下划线。
	Account = `^[a-zA-Z][a-zA-Z0-9_]{4,15}$`
	// Nickname matches nicknames with letters, digits, underscore, hyphen, or CJK.
	// Nickname 匹配昵称（字母、数字、下划线、连字符或中文）。
	Nickname = `[A-Za-z0-9_\-\u4e00-\u9fa5]+`
	// ChinaPhone1 matches mainland phone numbers like 0511-4405222 or 021-87888822.
	// ChinaPhone1 匹配国内电话号码（如 0511-4405222、021-87888822）。
	ChinaPhone1 = `\d{3}-\d{8}|\d{4}-\d{7}`
	// ChinaPhone2 matches mainland phone numbers (broader format).
	// ChinaPhone2 匹配国内电话号码（更宽格式）。
	ChinaPhone2 = `[0-9-()（）]{7,18}`
	// QQ matches Tencent QQ numbers.
	// QQ 匹配腾讯 QQ 号。
	QQ = `[1-9]([0-9]{5,11})`
	// ZipCode matches China postal codes (6 digits).
	// ZipCode 匹配中国邮政编码（6 位数字）。
	ZipCode = `[1-9]\d{5}(?!\d)`
	// ChinaIDCard matches China ID cards (15 or 18 digits).
	// ChinaIDCard 匹配中国身份证号（15 或 18 位）。
	ChinaIDCard = `\d{17}[\d|x]|\d{15}`
	// IPAddress matches IPv4 addresses.
	// IPAddress 匹配 IPv4 地址。
	IPAddress = `(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)`

	//-------------------------------------
	/*
	 匹配特定字符串：
	 ^[A-Za-z]+$　　//匹配由26个英文字母组成的字符串
	 ^[A-Z]+$　　//匹配由26个英文字母的大写组成的字符串
	 ^[a-z]+$　　//匹配由26个英文字母的小写组成的字符串
	 ^[A-Za-z0-9]+$　　//匹配由数字和26个英文字母组成的字符串
	 ^w+$　　//匹配由数字、26个英文字母或者下划线组成的字符串

	 在使用RegularExpressionValidator验证控件时的验证功能及其验证表达式介绍如下:
	 只能输入数字：“^[0-9]*$”
	 只能输入n位的数字：“^d{n}$”
	 只能输入至少n位数字：“^d{n,}$”
	 只能输入m-n位的数字：“^d{m,n}$”
	 只能输入零和非零开头的数字：“^(0|[1-9][0-9]*)$”
	 只能输入有两位小数的正实数：“^[0-9]+(.[0-9]{2})?$”
	 只能输入有1-3位小数的正实数：“^[0-9]+(.[0-9]{1,3})?$”
	 只能输入非零的正整数：“^+?[1-9][0-9]*$”
	 只能输入非零的负整数：“^-[1-9][0-9]*$”
	 只能输入长度为3的字符：“^.{3}$”
	 只能输入由26个英文字母组成的字符串：“^[A-Za-z]+$”
	 只能输入由26个大写英文字母组成的字符串：“^[A-Z]+$”
	 只能输入由26个小写英文字母组成的字符串：“^[a-z]+$”
	 只能输入由数字和26个英文字母组成的字符串：“^[A-Za-z0-9]+$”
	 只能输入由数字、26个英文字母或者下划线组成的字符串：“^w+$”
	*/
)

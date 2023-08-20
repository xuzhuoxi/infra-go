package regexpx

// 参考：http://en.wikipedia.org/wiki/Regular_expression
// POSIX正则表达式分为Basic Regular Expressions 和 Extended Regular Expressions。
// ERE增加支持?,+和|，去除了通配符()和{}。而且POSIX正则表达式的标准语法经常坚持使用附加的语法来支持特殊应用。虽然POSIX.2没有实现一些具体的细节，BRE和ERE提供被很多工具使用的标准。
// BRE要求通配符()和{}写成和\{\}，ERE中无需这样。

//数字相关
const (
	// Zero 零
	Zero = `^0$`
	// Digital 数字
	Digital = `^[0-9]$`

	// PInt 正整数
	PInt = `^[1-9]\d*$`
	// NInt 负整数
	NInt = `^-[1-9]\d*$`
	// NNInt 非负整数（正整数 + 0）
	NNInt = `^[1-9]\d*|0$`
	// NPInt 非正整数（负整数 + 0）
	NPInt = `^-[1-9]\d*|0$`
	// Int 整数
	Int = `^-?[1-9]\d*|0$`

	// PFloat 正浮点数
	PFloat = `^[1-9]\d*.\d*|0.\d*[1-9]\d*$`
	// NFloat 负浮点数
	NFloat = `^-([1-9]\d*.\d*|0.\d*[1-9]\d*)$`
	// NNFloat 非负浮点数（正浮点数 + 0）
	NNFloat = `^[1-9]\d*.\d*|0.\d*[1-9]\d*|0?.0+|0$`
	// NPFloat 非正浮点数（负浮点数 + 0）
	NPFloat = `^(-([1-9]\d*.\d*|0.\d*[1-9]\d*))|0?.0+|0$`
	// Float 浮点数
	Float = `^-?([1-9]\d*.d*|0.\d*[1-9]\d*|0?.0+|0)$`
)

// 常用
const (
	// LowerLetter 小写字母
	LowerLetter = `[a-z]`
	// UpperLetter 小写字母
	UpperLetter = `[A-Z]`
	// Letter 字母
	Letter = `[a-zA-Z]`
	// Space 空白
	Space = `[\t\n\f\r]`
	// Work 单词字符
	Work = `[0-9A-Za-z_]`

	// Chinese 中文字符 [u4e00-u9fa5]
	Chinese = `[\p{Han}]`
	// DoubleChar
	// 双字节字符
	// 包括汉字在内
	DoubleChar = `[^x00-xff]`
	// EmptyLine 空白行
	EmptyLine = `ns*r`
	// HTML
	// HTML标记
	// 上面这个也仅仅能匹配部分，对于复杂的嵌套标记依旧无能为力
	HTML = `<(S*?)[^>]*>.*?|<.*? /> `
	// EmptyHeadTail 首尾空白字符
	EmptyHeadTail = `^\s*|\s*$ `
	// EMail Email地址
	EMail = `w[-w.+]*@([A-Za-z0-9][-A-Za-z0-9]+.)+[A-Za-z]{2, 14}`
	// EMail2 Email地址
	EMail2 = `w+([-+.]w+)*@w+([-.]w+)*.w+([-.]w+)*`
	// URL 网址URL
	URL = `^((https|http|ftp|rtsp|mms)?://)[^s]+`
	// URL2 网址URL
	URL2 = `[a-zA-z]+://[^s]*`
	// Account 账号 字母开头，允许 5-16 个字节，允许字母数字下划线
	Account = `^[a-zA-Z][a-zA-Z0-9_]{4,15}$`
	// Nickname 昵称 字母开头，允许 5-16 个字节，允许字母数字下划线
	Nickname = `[A-Za-z0-9_\-\u4e00-\u9fa5]+`
	// ChinaPhone1 国内电话号码 匹配形式如 0511-4405222 或 021-87888822
	ChinaPhone1 = `\d{3}-\d{8}|\d{4}-\d{7}`
	// ChinaPhone2 国内电话号码
	ChinaPhone2 = `[0-9-()（）]{7,18}`
	// QQ 腾讯QQ号
	QQ = `[1-9]([0-9]{5,11})`
	// ZipCode 中国邮政编码 中国邮政编码为6位数字
	ZipCode = `[1-9]\d{5}(?!\d)`
	// ChinaIDCard 身份证 中国的身份证为15位或18位
	ChinaIDCard = `\d{17}[\d|x]|\d{15}`
	// IPAddress ip地址
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

package stringx

type PasswdFlag int

const (
	// N 0-9
	// Numbers 0-9
	// 数字 0-9
	N PasswdFlag = 1 << iota
	// L a-z
	// Lowercase letters a-z
	// 小写字母 a-z
	L
	// U
	// Uppercase letter A-Z
	// 大写字母 A-Z
	U
	// S
	// Symbols found on the keyboard (all keyboard characters not defined as letters or numerals) and spaces
	// 在键盘上找到的符号（所有未定义为字母或数字的键盘字符）和空格
	S
)

const (
	// LOrU
	// Uppercase or lowercase letter
	// 全部字母
	LOrU = L | U
	// DefaultPasswdFlag
	// Include Number and Letter
	// 数字和字母
	DefaultPasswdFlag            = N | LOrU
	mask              PasswdFlag = 0x11f
)

// CheckPassword
// Check that the password content is legitimate
// 检查密码内容是否合法
func CheckPassword(pwd string, flag PasswdFlag, minLen int, maxLen int) bool {
	if minLen > maxLen {
		minLen, maxLen = maxLen, minLen
	}
	pLen := len(pwd)
	if pLen > len([]rune(pwd)) { //包含多字符字符
		return false
	}
	if pLen < minLen || pLen > maxLen {
		return false
	}
	for _, c := range pwd {
		b := true
		switch {
		case c >= '0' && c <= '9':
			b = b && bitFit(flag, N)
		case c >= 'a' && c <= 'z':
			b = b && bitFit(flag, L)
		case c >= 'A' && c <= 'Z':
			b = b && bitFit(flag, U)
		case (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~'):
			b = b && bitFit(flag, S)
		default:
			return false
		}
		if !b {
			return false
		}
	}
	return true
}

func bitFit(value1, value2 PasswdFlag) bool {
	return value1&value2 > 0
}

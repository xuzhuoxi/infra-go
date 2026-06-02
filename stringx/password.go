package stringx

// PasswdFlag is a bitmask of allowed password character classes.
// PasswdFlag 表示密码允许的字符类别位掩码。
type PasswdFlag int

const (
	// N allows digits 0-9.
	// N 允许数字 0-9。
	N PasswdFlag = 1 << iota
	// L allows lowercase letters a-z.
	// L 允许小写字母 a-z。
	L
	// U allows uppercase letters A-Z.
	// U 允许大写字母 A-Z。
	U
	// S allows keyboard symbols and spaces.
	// S 允许键盘符号（非字母数字的键盘字符）和空格。
	S
)

const (
	// LOrU allows lowercase and uppercase letters.
	// LOrU 允许大小写字母。
	LOrU = L | U
	// DefaultPasswdFlag requires digits and letters.
	// DefaultPasswdFlag 默认要求数字和字母。
	DefaultPasswdFlag            = N | LOrU
	mask              PasswdFlag = 0x11f
)

// CheckPassword validates password length and allowed character classes.
// CheckPassword 检查密码长度及允许的字符类别是否合法。
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

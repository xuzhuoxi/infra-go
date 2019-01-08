package stringsx

type PasswdFlag int

const (
	// N Numbers
	N PasswdFlag = 1 << iota
	// L Lowercase letters
	L
	// U Uppercase letter
	U
	// S Symbols found on the keyboard (all keyboard characters not defined as letters or numerals) and spaces
	S
)

const (
	// L_OR_U Uppercase or lowercase letter
	L_OR_U = L | U
	// Include Number and Letter
	DefaultPasswdFlag            = N | L_OR_U
	mask              PasswdFlag = 0x11f
)

func PasswordCheck(pwd string, flag PasswdFlag, minLen int, maxLen int) bool {
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

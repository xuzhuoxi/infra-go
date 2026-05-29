package mathx

import (
	"strconv"
	"strings"
)

// SizeUnit represents a byte count using binary (1024-based) units.
// SizeUnit 表示以 1024 为基数的字节大小单位。
type SizeUnit uint64

const (
	// Byte is one byte.
	// Byte 为 1 字节。
	Byte SizeUnit = 1
	// KB is 1024 bytes.
	// KB 为 1024 字节。
	KB = 1024 * Byte
	// MB is 1024 kilobytes.
	// MB 为 1024 KB。
	MB = 1024 * KB
	// GB is 1024 megabytes.
	// GB 为 1024 MB。
	GB = 1024 * MB
	// TB is 1024 gigabytes.
	// TB 为 1024 GB。
	TB = 1024 * GB
	// PB is 1024 terabytes.
	// PB 为 1024 TB。
	PB = 1024 * TB
	// EB is 1024 petabytes.
	// EB 为 1024 PB。
	EB = 1024 * PB
)

// ParseSize parses a human-readable size string (e.g. "10MB", "1.5GB") into SizeUnit.
// ParseSize 将可读大小字符串（如 "10MB"、"1.5GB"）解析为 SizeUnit。
func ParseSize(sizeStr string) SizeUnit {
	sizeStr = strings.ToUpper(strings.TrimSpace(sizeStr))
	if len(sizeStr) == 0 {
		return 0
	}
	var size SizeUnit
	var err error
	if strings.HasSuffix(sizeStr, "KB") {
		size, err = parseSize(sizeStr, KB)
	} else if strings.HasSuffix(sizeStr, "MB") {
		size, err = parseSize(sizeStr, MB)
	} else if strings.HasSuffix(sizeStr, "GB") {
		size, err = parseSize(sizeStr, GB)
	} else if strings.HasSuffix(sizeStr, "TB") {
		size, err = parseSize(sizeStr, TB)
	} else if strings.HasSuffix(sizeStr, "PB") {
		size, err = parseSize(sizeStr, PB)
	} else if strings.HasSuffix(sizeStr, "EB") {
		size, err = parseSize(sizeStr, EB)
	} else {
		value, err := strconv.ParseFloat(strings.TrimSpace(sizeStr), 64)
		if err != nil {
			return 0
		}
		return SizeUnit(value)
	}
	if err != nil {
		return 0
	}
	return size
}

func parseSize(sizeStr string, unit SizeUnit) (size SizeUnit, err error) {
	sizeStr = sizeStr[:len(sizeStr)-2]
	sizeStr = strings.TrimSpace(sizeStr)
	value, err1 := strconv.ParseFloat(sizeStr, 64)
	if nil != err1 {
		return 0, err1
	}
	return SizeUnit(value * float64(unit)), nil
}

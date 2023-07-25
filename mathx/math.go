package mathx

import (
	"strconv"
	"strings"
)

type SizeUnit uint64

const (
	Byte SizeUnit = 1
	KB            = 1024 * Byte
	MB            = 1024 * KB
	GB            = 1024 * MB
	TB            = 1024 * GB
	PB            = 1024 * TB
	EB            = 1024 * PB
)

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

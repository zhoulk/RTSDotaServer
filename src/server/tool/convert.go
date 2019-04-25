package tool

import (
	"regexp"
	"strconv"
	"strings"
)

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func IsNum(str string) bool {
	pattern := "\\d+" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, str)
	return result
}

func JoinInt32Arr(arr []int32, seperator string) string {
	strs := make([]string, 0)
	for _, i := range arr {
		strs = append(strs, strconv.FormatInt(int64(i), 10))
	}
	return strings.Join(strs, seperator)
}

func SpliteInt32Arr(str string, seperator string) []int32 {
	strings := strings.Split(str, seperator)
	arr := make([]int32, 0)
	for _, s := range strings {
		i, err := strconv.ParseInt(s, 0, 32)
		if err == nil {
			arr = append(arr, int32(i))
		}
	}
	return arr
}

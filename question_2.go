package stream_test

import (
	"strings"

	"github.com/hrz123/stream"
	"github.com/hrz123/stream/types"
)

// - Q1: 计算一个 string 中小写字母的个数
func Question2Sub1(str string) int64 {
	return stream.OfSlice(strings.Split(str, "")).Filter(func(e types.T) bool {
		return "a" <= e.(string) && e.(string) <= "z"
	}).Count()
}

// - Q2: 找出 []string 中，包含小写字母最多的字符串
func Question2Sub2(list []string) string {
	var maxCount int64
	var s string
	stream.OfSlice(list).ForEach(func(e types.T) {
		if count := Question2Sub1(e.(string)); count > maxCount {
			s = e.(string)
			maxCount = count
		}
	})
	return s
}

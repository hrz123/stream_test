package stream_test

import (
	"strings"

	"github.com/hrz123/stream"
	"github.com/hrz123/stream/types"
)

// - Q1: 计算一个 string 中小写字母的个数
func Question2Sub1(str string) int64 {
	stream.OfSlice(strings.Split(str, "")).Filter(func(e types.T) bool {
		return "a" <= e.(string) && e.(string) <= "z"
	})
	return 0
}

// - Q2: 找出 []string 中，包含小写字母最多的字符串
func Question2Sub2(list []string) string {
	return (stream.OfSlice(list).Sorted(func(left types.T, right types.T) int {
		return int(Question2Sub1(left.(string)) - Question2Sub1(right.(string)))
	}).Limit(1).ToElementSlice("")).([]string)[0]
}

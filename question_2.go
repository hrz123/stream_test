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
	return stream.OfSlice(list).
		Map(func(e types.T) types.R {
			return &types.Pair{
				First:  e,
				Second: Question2Sub1(e.(string)),
			}
		}).Reduce(func(e1 types.T, e2 types.T) types.T {
		if e1.(*types.Pair).Second.(int64) > e2.(*types.Pair).Second.(int64) {
			return e1
		}
		return e2
	}).Get().(*types.Pair).First.(string)
}

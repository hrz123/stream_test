package stream_test

import (
	"strings"

	"github.com/hrz123/stream"
	"github.com/hrz123/stream/types"
)

// - Q1: 输入一个整数 int，字符串string。将这个字符串重复n遍返回
func Question3Sub1(str string, n int) string {
	builder := stream.RepeatN(str, int64(n)).ReduceWith(strings.Builder{}, func(acc types.R, e types.T) types.R {
		builder := acc.(strings.Builder)
		builder.WriteString(e.(string))
		return builder
	}).(strings.Builder)
	return builder.String()
}

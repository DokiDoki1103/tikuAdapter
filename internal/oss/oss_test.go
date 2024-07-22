package oss

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Sort(t *testing.T) {
	var answer = make([]string, 0)
	answer = append(answer, "A")
	answer = append(answer, "C")
	answer = append(answer, "B")

	// 创建 answer 的一个副本
	noSortAnswer := make([]string, len(answer))
	copy(noSortAnswer, answer)

	fmt.Println(answer)

	sort.Strings(answer)
	fmt.Println(answer)
	fmt.Println(noSortAnswer)
}

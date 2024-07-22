package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sort"
	"testing"
)

func Test_OSS(t *testing.T) {
	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "", "")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("1:")

	bucket, err := client.Bucket("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = bucket.PutObjectFromFile("", "")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

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

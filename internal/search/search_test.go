package search

import (
	"fmt"
	"github.com/itihey/tikuAdapter/internal/model"
	"testing"
)

func TestSearchWanneng(t *testing.T) {
	var client = searchWannengClient{}
	testRequest := model.SearchRequest{
		Question: "中国最美丽的风景",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
	fmt.Println("成功搜题", response.Answer)
}

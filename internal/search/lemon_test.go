package search

import (
	"encoding/json"
	"fmt"
	"github.com/itihey/tikuAdapter/pkg/model"
	"testing"
)

// TestSearchLemonClient_SearchAnswer 测试柠檬题库接口
func TestSearchLemonClient_SearchAnswer(t *testing.T) {
	var client = LemonClient{
		Disable: false,
		Token:   "8a3debe92e2ba83d6786e186bef2a424",
	}

	testRequest := model.SearchRequest{
		Question: "大学生缓解性欲的常见途径是()",
	}

	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求柠檬题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试柠檬题库 成功搜题", string(marshal))
}

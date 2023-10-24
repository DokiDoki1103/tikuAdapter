package search

import (
	"encoding/json"
	"fmt"
	"github.com/itihey/tikuAdapter/internal/model"
	"testing"
)

func TestSearchIcodefClient_SearchAnswer(t *testing.T) {
	var client = SearchIcodefClient{}
	testRequest := model.SearchRequest{
		Question: "下面选项中,属于男性在青春期生理变化的内容有?()",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Println("TestSearchIcodefClient_SearchAnswer 测试失败 只是对方的接口禁止了国外访问，故直接pass")
		// t.Errorf("Expected no error, but got an error: %v", err)
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("TestSearchIcodefClient_SearchAnswer 成功搜题", string(marshal))
}
func TestSearchWanneng(t *testing.T) {
	var client = SearchClient{}
	testRequest := model.SearchRequest{
		Question: "下面选项中,属于男性在青春期生理变化的内容有?()",
	}

	// 调用被测试的方法
	response, err := client.Wanneng.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("Expected no error, but got an error: %v", err)
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("TestSearchWanneng 成功搜题", string(marshal))
}

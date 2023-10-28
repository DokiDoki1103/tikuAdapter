package search

import (
	"encoding/json"
	"fmt"
	"github.com/itihey/tikuAdapter/pkg/model"
	"os"
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
		t.Errorf("请求icodef题库异常: %v", err)
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试icodef题库 成功搜题", string(marshal))
}

func TestSearchWannengClient_SearchAnswer(t *testing.T) {
	var client = SearchWannengClient{
		Disable: false,
	}
	testRequest := model.SearchRequest{
		Question: "下面选项中,属于男性在青春期生理变化的内容有?()",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("请求万能题库异常: %v", err)
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试万能题库 成功搜题", string(marshal))
}

func TestSearchEnncyClient_SearchAnswer(t *testing.T) {
	var client = SearchEnncyClient{
		Disable: false,
		Token:   os.Getenv("ENNCY_TOKEN"),
	}

	testRequest := model.SearchRequest{
		Question: "《红楼梦》传世抄本中，惟一一个可以确定抄录年代的，原为吴晓玲先生旧藏、今藏于首都图书馆的是？",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("请求enncy题库异常: %v", err)
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试enncy题库 成功搜题", string(marshal))
}

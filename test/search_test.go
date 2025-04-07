package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/itihey/tikuAdapter/internal/search"
	"github.com/itihey/tikuAdapter/pkg/model"
)

func TestSearchIcodefClient_SearchAnswer(t *testing.T) {
	var client = search.IcodefClient{}
	testRequest := model.SearchRequest{
		Question: "下面选项中,属于男性在青春期生理变化的内容有?()",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("请求icodef题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试icodef题库 成功搜题", string(marshal))
}

func TestSearchWannengClient_SearchAnswer(t *testing.T) {
	var client = search.WannengClient{
		Enable: true,
	}
	testRequest := model.SearchRequest{
		Question: "大学生缓解性欲的常见途径是()",
		Options:  []string{"自然缓解", "升华", "转移", "克制"},
		Type:     0,
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求万能题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试万能题库 成功搜题", string(marshal))
}

func TestSearchEnncyClient_SearchAnswer(t *testing.T) {
	var client = search.EnncyClient{
		Enable: true,
		Token:  os.Getenv("ENNCY_TOKEN"),
	}

	testRequest := model.SearchRequest{
		Question: "《红楼梦》传世抄本中，惟一一个可以确定抄录年代的，原为吴晓玲先生旧藏、今藏于首都图书馆的是？",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求enncy题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试enncy题库 成功搜题", string(marshal))
}

func TestBuguake_SearchAnswer(t *testing.T) {
	var client = search.BuguakeClient{
		Enable: true,
	}

	testRequest := model.SearchRequest{
		Question: "苏轼的《西江月》是其在哪里所创作的?( )",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("请求不挂科题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试不挂科题库 成功搜题", string(marshal))
}

func TestAidian_SearchAnswer(t *testing.T) {

	var client = search.AidianClient{
		Enable: true,
	}

	testRequest := model.SearchRequest{
		Question: "9.十进制数(9)10比十六进制数(9)16小.( )",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求爱点题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试爱点题库 成功搜题", string(marshal))
}

// TestSearchLemonClient_SearchAnswer 测试柠檬题库接口
func TestSearchLemonClient_SearchAnswer(t *testing.T) {

	var client = search.LemonClient{
		Enable: true,
		Token:  "8a3debe92e2ba83d6786e186bef2a424",
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

// TestSearchTikuhaiClient_SearchAnswer 测试题库海题库接口
func TestSearchTikuhaiClient_SearchAnswer(t *testing.T) {
	var client = search.TikuhaiClient{
		Enable: true,
		Token:  "",
	}
	testRequest := model.SearchRequest{
		Question: "毛泽东思想形成的时代背景是( )",
		Options:  []string{"帝国主义战争与无产阶级革命成为时代主题", "和平与发展成为时代主题", "世界多极化成为时代主题", "经济全球化成为时代主题"},
		Type:     0,
	}
	//这道题付费题库里才有答案
	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求题库海题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试题库海题库 成功搜题", string(marshal))
}

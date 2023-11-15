package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/middleware"
	"github.com/itihey/tikuAdapter/internal/search"
	"github.com/itihey/tikuAdapter/pkg/global"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/itihey/tikuAdapter/pkg/util"
	"net/http"
	"reflect"
	"sync"
)

// Search 搜题接口
func Search(c *gin.Context) {
	var req model.SearchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, global.ErrorParam)
		return
	}

	var result [][]string               // 最后所有的答案的二维数组
	var answer [][]string               // 本地答案
	if c.Query("localDisable") != "1" { // 没有禁用本地搜索的话，先查询本地
		answer, err = search.GetDBSearch().SearchAnswer(req)
		if err != nil {
			logger.SysError(fmt.Sprintf("查询本地答案出错：%s", err.Error()))
		}
		result = append(result, answer...)
	}

	// 再查询第三方
	if len(result) == 0 {
		searchClient := search.Client{
			Wanneng: &search.WannengClient{
				Token:   c.Query("wannengToken"),
				Disable: c.Query("wannengDisable") == "1",
			},
			Icodef: &search.IcodefClient{
				Token:   c.Query("icodefToken"),
				Disable: c.Query("icodefDisable") == "1",
			},
			Enncy: &search.EnncyClient{
				Token:   c.Query("enncyToken"),
				Disable: c.Query("enncyDisable") == "1",
			},
			Buguake: &search.BuguakeClient{
				Disable: c.Query("buguakeDisable") == "1",
			},
			AiDian: &search.AidianClient{
				Disable: true,
				YToken:  c.Query("aidianYToken"),
			},
		}
		var wg sync.WaitGroup
		var mu sync.Mutex
		val := reflect.ValueOf(&searchClient).Elem()

		for i := 0; i < val.NumField(); i++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				clientField := val.Field(idx)
				methodValue := clientField.MethodByName("SearchAnswer")
				requestValue := reflect.ValueOf(req)
				res := methodValue.Call([]reflect.Value{requestValue})

				if len(res) > 1 && !res[1].IsNil() { // 出现错误
					e := res[1].Interface().(error).Error()
					logger.SysError(fmt.Sprintf("调用%s接口出错：%s", clientField.Type().String(), e))
				} else {
					mu.Lock()
					defer mu.Unlock()
					ans := res[0].Interface().([][]string)
					result = append(result, ans...)
				}

			}(i)
		}
		wg.Wait()
	}

	if len(result) > 0 {
		resp := util.FillAnswerResponse(result, &req)
		if len(answer) == 0 && c.Query("localDisable") != "1" {
			middleware.CollectAnswer(resp)
		}
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusNotFound, global.ErrorQuestionNotFound)
	}
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/search"
	"github.com/itihey/tikuAdapter/pkg/global"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/itihey/tikuAdapter/pkg/util"
	"net/http"
	"reflect"
	"sync"
)

func Search(c *gin.Context) {
	searchClient := search.SearchClient{
		Wanneng: &search.SearchWannengClient{
			Token:   c.Query("wannengToken"),
			Disable: c.Query("wannengDisable") == "1",
		},
		Icodef: &search.SearchIcodefClient{
			Token:   c.Query("icodefToken"),
			Disable: c.Query("icodefDisable") == "1",
		},
		Enncy: &search.SearchEnncyClient{
			Token:   c.Query("enncyToken"),
			Disable: c.Query("enncyDisable") == "1",
		},
	}

	var req model.SearchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var result [][]string // 最后所有的答案的二维数组
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

	if len(result) > 0 {
		resp := util.FillAnswerResponse(result, &req)
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusNotFound, global.ErrorQuestionNotFound)
	}
}

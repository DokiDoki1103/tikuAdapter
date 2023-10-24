package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/model"
	"github.com/itihey/tikuAdapter/internal/search"
	"github.com/itihey/tikuAdapter/pkg/global"
	"net/http"
	"reflect"
	"sync"
)

func Search(c *gin.Context) {
	searchClient := search.SearchClient{}

	var req model.SearchRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var result [][]string
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

			if len(res) > 1 && !res[1].IsNil() {
				fmt.Println(res[1].Interface().(error).Error())
			} else {
				mu.Lock()
				defer mu.Unlock()
				ans := res[0].Interface().([][]string)
				result = append(result, ans...)
			}

		}(i)
	}
	wg.Wait()

	fmt.Println()
	resp := model.SearchResponse{
		MoreAnswer: result,
		Question:   req.Question,
		Options:    req.Options,
		Type:       req.Type,
		Plat:       req.Plat,
	}
	if len(result) > 0 {
		resp.Answer = result[0]
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusNotFound, global.ErrorQuestionNotFound)
	}
}

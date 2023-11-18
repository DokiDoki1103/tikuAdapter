package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itihey/tikuAdapter/internal/middleware"
	"github.com/itihey/tikuAdapter/internal/registry/manager"
	"github.com/itihey/tikuAdapter/internal/search"
	"github.com/itihey/tikuAdapter/pkg/global"
	"github.com/itihey/tikuAdapter/pkg/logger"
	"github.com/itihey/tikuAdapter/pkg/model"
	"github.com/itihey/tikuAdapter/pkg/util"
	"net/http"
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
		var clients = []search.Search{
			&search.WannengClient{
				Disable: c.Query("wannengDisable") == "1",
			},
			&search.IcodefClient{
				Token:   c.Query("icodefToken"),
				Disable: c.Query("icodefDisable") == "1",
			},
			&search.EnncyClient{
				Token:   c.Query("enncyToken"),
				Disable: c.Query("enncyDisable") == "1",
			},
			&search.BuguakeClient{
				Disable: c.Query("buguakeDisable") == "1",
			},
			&search.AidianClient{
				Disable: c.Query("aidianDisable") == "1",
				YToken:  c.Query("aidianYToken"),
			},
		}
		cfg := manager.GetManager().GetConfig()
		for i := range cfg.API {
			clients = append(clients, cfg.API[i])
		}

		var wg sync.WaitGroup
		var mu sync.Mutex

		for i := range clients {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				res, err := clients[idx].SearchAnswer(req)
				if err == nil && len(res) > 0 {
					mu.Lock()
					defer mu.Unlock()
					result = append(result, res...)
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

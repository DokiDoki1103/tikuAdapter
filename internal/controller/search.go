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
	"strings"
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

	var result [][]string                          // 最后所有的答案的二维数组
	var localAnswer [][]string                     // 本地答案
	if strings.Contains(c.Query("use"), "local") { // 使用本地题库的话
		localAnswer, err = search.GetDBSearch().SearchAnswer(req)
		if err != nil {
			logger.SysError(fmt.Sprintf("查询本地答案出错：%s", err.Error()))
		}
		result = append(result, localAnswer...)
	}

	// 再查询第三方
	if len(result) == 0 {
		var clients = []search.Search{
			&search.BuguakeClient{
				Enable: strings.Contains(c.Query("use"), "buguake") || c.Query("use") == "",
			},
			&search.IcodefClient{
				Token:  c.Query("icodefToken"),
				Enable: strings.Contains(c.Query("use"), "icodef") || c.Query("use") == "",
			},
			&search.WannengClient{
				Token:  c.Query("wannengToken"),
				Enable: strings.Contains(c.Query("use"), "wanneng") || c.Query("use") == "",
			},
			&search.EnncyClient{
				Token:  c.Query("enncyToken"),
				Enable: strings.Contains(c.Query("use"), "enncy"),
			},
			&search.AidianClient{
				Enable: strings.Contains(c.Query("use"), "aidian"),
				YToken: c.Query("aidianYToken"),
			},
			&search.LemonClient{
				Enable: strings.Contains(c.Query("use"), "lemon"),
				Token:  c.Query("lemonToken"),
			},
		}
		cfg := manager.GetManager().GetConfig()
		for _, api := range cfg.API {
			if strings.Contains(c.Query("use"), api.Name) {
				clients = append(clients, api)
			}
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

	resp := util.FillAnswerResponse(result, &req)

	if len(localAnswer) == 0 && c.Query("noRecord") == "" { // 只有本地题库没有答案或者明确标志不要记录
		middleware.CollectAnswer(resp, c.Query("extra"))
	}
	c.JSON(http.StatusOK, resp)
}

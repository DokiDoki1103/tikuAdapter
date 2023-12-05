package timer

import (
	"github.com/itihey/tikuAdapter/internal/registry/manager"
	"github.com/itihey/tikuAdapter/internal/service"
	"github.com/robfig/cron/v3"
)

// StartTimer 启动定时器
func StartTimer() {
	if len(manager.GetManager().GetConfig().Elasticsearch.Addresses) > 0 {
		// cron格式（秒，分，时，天，月，周）
		c := cron.New(cron.WithSeconds())
		c.AddFunc("*/30 * * * * *", service.SyncElasticsearch)
		c.Start()
	}
}

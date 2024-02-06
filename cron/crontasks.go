package cron

import (
	"time"

	"github.com/robfig/cron/v3"
	"oplesko.com/tft_pipeline/database"
	"oplesko.com/tft_pipeline/pipelines/rankedplayers"
)

func RunCronTasks() {
	c := cron.New()

	c.AddFunc("@weekly", rankedplayers.BeginRankedPlayerPipeline)
	c.AddFunc("@daily", func() { database.PruneMatchIds(4 * 24 * time.Hour) })

	c.Run()
}

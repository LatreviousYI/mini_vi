package common

import (
	"log"

	"github.com/go-co-op/gocron/v2"
)

var GoCronTab gocron.Scheduler

func init() {
	var err error
	GoCronTab, err = gocron.NewScheduler()
	if err != nil {
		log.Printf("定时任务启动报错%s", err)
	}
	GoCronTab.Start()
}

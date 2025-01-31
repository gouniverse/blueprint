package scheduler

import (
	"project/internal/tasks"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/mingrammer/cfmt"
)

// StartAsync starts the scheduler in the backgroun without blocking the main thread
func StartAsync() {
	scheduler := gocron.NewScheduler(time.UTC)

	// Example of task scheduled every 2 minutes
	// only on production and staging, not on dev and local
	// if config.IsEnvStaging() || config.IsEnvProduction() {
	// 	scheduler.Every(2).Minutes().Do(func() {
	// 		_, err := taskhandlers.NewHelloWorldTaskHandler().Enqueue()
	// 		if err != nil {
	// 			cfmt.Errorln(err.Error())
	// 		}
	// 	})
	// }

	// Example of daily scheduled task
	// scheduler.Every(1).Day().At("01:00").Do(func() {
	// 	_, err := taskhandlers.NewHelloWorldTaskHandler().Enqueue()
	// 	if err != nil {
	// 		cfmt.Errorln(err.Error())
	// 	}
	// })

	// Schedule Building the Cache Every 2 Minutes
	// only on production, no need on dev and local
	// if config.IsEnvStaging() || config.IsEnvProduction() {
	// 	scheduler.Every(2).Minutes().Do(func() {
	// 		pool.BuildCache()
	// 	})
	// }

	// Schedule Building the Stats Every 2 Minutes
	scheduler.Every(2).Minutes().Do(func() {
		_, err := tasks.NewStatsVisitorEnhanceTask().Enqueue()
		if err != nil {
			cfmt.Errorln(err.Error())
		}
	})

	scheduler.StartAsync()
}

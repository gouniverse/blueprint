package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
)

func Initialize() {
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

	scheduler.StartAsync()
}

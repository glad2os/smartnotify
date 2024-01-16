package utils

import (
	"github.com/robfig/cron/v3"
	"time"
)

func GetNextRunTime(cronExpr string) (time.Time, error) {
	scheduler := cron.New()
	// AddFunc only to get the next run time
	_, err := scheduler.AddFunc(cronExpr, func() {})
	if err != nil {
		return time.Time{}, err
	}
	scheduler.Start()
	defer scheduler.Stop()

	return scheduler.Entries()[0].Next, nil
}

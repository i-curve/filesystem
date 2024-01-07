package handle

import (
	"time"
)

type Cron struct {
	ID         int       `json:"id"`
	Bucket     string    `json:"bucket"`
	Path       string    `json:"path"`
	DeleteTime time.Time `json:"delete_time"`
}

func CronDelete(cron *Cron) {
	time.AfterFunc(-time.Since(cron.DeleteTime), func() {
		removeFile(cron.Bucket, cron.Path)
		mariadb.Delete(&cron)
	})
}

package handle

import (
	"testing"
	"time"
)

func TestDeleteCron(t *testing.T) {
	CronDelete(&Cron{DeleteTime: time.Now().Add(time.Second * 2 * -1)})
	// time.Sleep(3 * time.Second)
}

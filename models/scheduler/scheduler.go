/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         scheduler.go
@ Create Time:  2020/5/6 16:02
@ Software:     GoLand
*/

package scheduler

import (
	"github.com/robfig/cron/v3"
)

var Scheduler *cron.Cron
var DelChannel chan int

func init() {
	Scheduler = cron.New()
	Scheduler.Start()

	DelChannel = make(chan int, 32)
	go func() {
		for {
			select {
			case id := <-DelChannel:
				Scheduler.Remove(cron.EntryID(id))
			}
		}
	}()
}

package timetrigger

import (
	"time"
)

// HandlerFunc 处理函数类型
type HandlerFunc func()

// Trigger 时间触发器，根据给定的时间触发函数，每天一次
type Trigger struct {
	// 给定小时
	hour int
	// 给定分钟
	min int
	// 触发函数
	trigerfunc HandlerFunc
}

// New 生成新的Trigger
func New(hour int, min int, trigerfunc HandlerFunc) *Trigger {
	return &Trigger{hour, min, trigerfunc}
}

// 启动触发器
func (trigger *Trigger) Run() {
	nowtime := time.Now()
	go func() {
		d, _ := time.ParseDuration("24h")
		tm := time.Date(nowtime.Year(), nowtime.Month(), nowtime.Day(), trigger.hour, trigger.min, 0, 0, nowtime.Location())
		if nowtime.After(tm) {
			tm = tm.Add(d)
		}
		for {
			// sleep for time
			time.Sleep(tm.Sub(time.Now()))
			trigger.trigerfunc()
			tm = tm.Add(d)
		}
	}()
}

package timetrigger

import (
	"time"
)

// Trigger 时间触发器，根据给定的时间触发函数，每天一次
type Trigger struct {
	// 开始小时
	tm time.Time
	// 延迟时间（小时）
	interval time.Duration
	// 触发函数
	triggerfunc func()
	// 控制器
	timer *time.Timer
}

// New 生成新的Trigger
func New(tm time.Time, interval time.Duration, triggerfunc func()) *Trigger {
	trigger := &Trigger{}
	trigger.tm = tm
	trigger.interval = interval
	trigger.triggerfunc = triggerfunc
	return trigger
}

// Run 启动触发器
func (trigger *Trigger) Run() {
	for nowtime := time.Now(); nowtime.After(trigger.tm); trigger.tm = trigger.tm.Add(trigger.interval) {
	}
	// sleep for time
	trigger.timer = time.AfterFunc(trigger.tm.Sub(time.Now()), func() {
		ticker := time.NewTicker(trigger.interval)
		for {
			trigger.triggerfunc()
			<-ticker.C
		}
	})
}

// Stop 停止触发器
func (trigger *Trigger) Stop() {
	trigger.timer.Stop()
}

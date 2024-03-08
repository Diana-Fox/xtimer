package domain

import (
	"github.com/robfig/cron/v3"
	"time"
)

type MethodType int

// 不允许定时任务有别的请求类型
const (
	POST MethodType = 1
	GET  MethodType = 2
)

type Job struct {
	Id        int64
	Name      string //注册的任务名
	Cron      string
	Method    MethodType //请求方法
	Parameter string     //请求参数
	Url       string     //请求参数
	JobType   string
	Version   int
}

var parser = cron.NewParser(cron.Minute | cron.Hour | cron.Dom |
	cron.Month | cron.Dow | cron.Descriptor)

func (j Job) NextTime() time.Time {
	//下一次 利用cron算
	parse, _ := parser.Parse(j.Cron)
	return parse.Next(time.Now())
}

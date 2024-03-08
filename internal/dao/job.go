package dao

import (
	"container/list"
	"context"
)

type JobDao interface {
	Preempt(ctx context.Context) *list.List
}
type Job struct {
	Id        int64
	Name      string //注册的任务名
	Cron      string
	Method    int    //请求方法
	Parameter string //请求参数
	Url       string //请求参数
	JobType   string
	Version   int
}
type CronJobDao struct {
}

func (c *CronJobDao) Preempt(ctx context.Context) *list.List {
	//TODO implement me
	panic("implement me")
}

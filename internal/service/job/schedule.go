package job

import (
	"context"
	domain "github.com/Diana-Fox/xtimer/internal/domian"
	"github.com/Diana-Fox/xtimer/internal/repository"
	"github.com/Diana-Fox/xtimer/internal/utils"
)

type ScheduleService interface {
	//抢占任务
	Preempt(ctx context.Context) (domain.Job, error)
	//重置下一次执行的时间
	ReleaseJob(ctx context.Context, job domain.Job) error
}
type ScheduleJobService struct {
	l utils.Logger //打日志
	r repository.JobRepository
}

// Preempt 去抢占一个任务
func (c *ScheduleJobService) Preempt(ctx context.Context) (domain.Job, error) {
	return c.r.Preempt(ctx)
}

// 释放本次调度状态
func (c *ScheduleJobService) ReleaseJob(ctx context.Context, job domain.Job) error {
	//TODO 根据表达式去计算
	//next := job.NextTime()
	//然后一次性把下次执行时间和状态都改掉
	panic("implement me")
}

package job

import (
	"context"
	domain "github.com/Diana-Fox/xtimer/internal/domian"
	"github.com/Diana-Fox/xtimer/internal/utils"
)

type JobService interface {
	//注册定时任务
	Register(ctx context.Context, job domain.Job) error
	//开启任务或者直接关闭任务
	Status(ctx context.Context, id int64, status int) error
}
type CronJobService struct {
	l utils.Logger //打日志
}

func (c *CronJobService) Register(ctx context.Context, job domain.Job) error {
	//TODO implement me
	panic("implement me")
}

func (c CronJobService) Status(ctx context.Context, id int64, status int) error {
	//TODO implement me
	panic("implement me")
}

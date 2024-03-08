package repository

import (
	"container/list"
	"context"
	"errors"
	"github.com/Diana-Fox/xtimer/internal/dao"
	domain "github.com/Diana-Fox/xtimer/internal/domian"
	"github.com/Diana-Fox/xtimer/internal/utils"
	"sync"
	"time"
)

var (
	LIST_NOT_JOB = errors.New("集合里面没有任务了")
	JOBS         = list.New()
	mx           = sync.Mutex{}
)

type JobRepository interface {
	//注册定时任务
	Register(ctx context.Context, job domain.Job) error
	//抢占任务
	Preempt(ctx context.Context) (domain.Job, error)
	//释放任务
	Release(ctx context.Context, id int64, version int) error
	//更新活跃时间，为了执行期间保活
	UpdateUtime(ctx context.Context, id int64, version int) error
	//更新下一次
	UpdateNextUtime(ctx context.Context, id int64, version int, next time.Time) error
	//修改任务状态
	Status(ctx context.Context, id int64, status int) error
}

// 任务
type CronJobRepository struct {
	l    utils.Logger //打日志
	dao  dao.JobDao
	jobs []dao.Job
}

func (c *CronJobRepository) Register(ctx context.Context, job domain.Job) error {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobRepository) Preempt(ctx context.Context) (domain.Job, error) {
	mx.Lock()
	mx.Unlock()
	if JOBS.Len() > 1 {
		//去取值
		job := JOBS.Front().Value.(dao.Job)
		return c.dao2domain(job), nil
	}
	//说明没有任务了，去数据库拿
	for {
		JOBS = c.dao.Preempt(ctx)
		if JOBS.Len() > 1 {
			//能到这，说明有了
			job := JOBS.Front().Value.(dao.Job)
			return c.dao2domain(job), nil
		}
		//说明没查到，在这歇歇
		time.Sleep(time.Second)
	}
}
func (c *CronJobRepository) dao2domain(d dao.Job) domain.Job {
	return domain.Job{
		Id:        d.Id,
		Name:      d.Name,
		Cron:      d.Cron,
		Method:    domain.MethodType(d.Method),
		Parameter: d.Parameter,
		Url:       d.Url,
		JobType:   d.JobType,
		Version:   d.Version,
	}
}

//	func (c *CronJobRepository) fistJob() (domain.Job, error) {
//		//todo 这里需要上个锁，一次只能一个进入
//		if len(c.jobs) == 0 {
//			return domain.Job{}, LIST_NOT_JOB
//		}
//		job := c.jobs[0]
//		c.jobs = c.jobs[1:]
//		return domain.Job{
//			Id:        job.Id,
//			Name:      job.Name,
//			Cron:      job.Cron,
//			Method:    domain.MethodType(job.Method),
//			Parameter: job.Parameter,
//			Url:       job.Url,
//			JobType:   job.JobType,
//			Version:   job.Version,
//		}, nil
//	}
func (c *CronJobRepository) Release(ctx context.Context, id int64, version int) error {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobRepository) UpdateUtime(ctx context.Context, id int64, version int) error {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobRepository) UpdateNextUtime(ctx context.Context, id int64, version int, next time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (c *CronJobRepository) Status(ctx context.Context, id int64, status int) error {
	//TODO implement me
	panic("implement me")
}

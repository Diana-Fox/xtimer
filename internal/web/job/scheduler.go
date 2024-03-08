package job

import (
	"context"
	"github.com/Diana-Fox/xtimer/internal/service/job"
	"github.com/Diana-Fox/xtimer/internal/thirdparty"
	"github.com/Diana-Fox/xtimer/internal/utils"
	"time"
)

// 调度器
type Scheduler struct {
	svc job.ScheduleJobService //获取任务信息
	exe thirdparty.Executor    //执行任务信息
	l   utils.Logger           //打日志
}

func (s *Scheduler) Schedule(ctx context.Context) error {
	for {
		println("调度器运行")
		////一次调度的数据库查询时间
		dbCtx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		//todo 需要想办法限制抢占流量
		job, err := s.svc.Preempt(dbCtx)
		if err != nil {
			//继续下一轮
			s.l.ERROR("抢占任务失败:%s", err.Error())
			continue
		}
		//执行job，去发起http请求,这里后面考虑替换成有失败策略的
		err = s.exe.Execute(ctx, job.Method,
			job.Parameter, job.Url, nil)
		if err != nil {
			//失败了，不需要释放，直接去下一轮
			continue
		}
		//顺利执行完成了，需要去执行释放
		go func() {
			ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
			defer cancel1()
			s.svc.ReleaseJob(ctx1, job) //去释放
		}()
	}
}

package main

import (
	"context"
	"time"

	"github.com/reugn/go-quartz/quartz"
)

func main() {
	ctx := context.Background()
	sched := quartz.NewStdScheduler()
	sched.Start(ctx)

	functionJob := quartz.NewFunctionJob(func(_ context.Context) (int, error) { return 42, nil })
	sched.ScheduleJob(functionJob, quartz.NewSimpleTrigger(time.Second*5))
	sched.Stop()
	sched.Wait(ctx)

}

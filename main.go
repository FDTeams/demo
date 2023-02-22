package main

import (
	"context"
	"net/http"
	"time"

	"github.com/seaung/gexample/demo/go-quartz/quartz"
)

func main() {
	ctx := context.Background()
	sched := quartz.NewStdScheduler()
	sched.Start(ctx)
	cronTrigger, _ := quartz.NewCronTrigger("1/5 * * * * *")
	shellJob := quartz.NewShellJob("ls -la")
	curlJob, _ := quartz.NewCurlJob(http.MethodGet, "http://worldclockapi.com/api/json/est/now", "", nil)
	functionJob := quartz.NewFunctionJob(func(_ context.Context) (int, error) { return 42, nil })
	sched.ScheduleJob(shellJob, cronTrigger)
	sched.ScheduleJob(curlJob, quartz.NewSimpleTrigger(time.Second*7))
	sched.ScheduleJob(functionJob, quartz.NewSimpleTrigger(time.Second*5))
	sched.Stop()
	sched.Wait(ctx)
}

package main
import (
    task"github.com/timer/task"
    "github.com/timer/service"
)
func main (){
  startTimer()
}

func startTimer()  {
	var tasks []*task.TimeTask
	tasks = append(tasks, demoTimerTask())
	cfg.StartSchedule(tasks, false)
}

var demoTimer *task.TimeTask
func demoTimerTask() *task.TimeTask {
	if demoTimer == nil {
		demoTimer = &cfg.SwTask{
			Name:"updateRedisDay",
			Run:service.UpdateDays, // 具体的业务逻辑
			Spec:"0 0 24 * * ?", // 触发时间
		}
	}
	return demoTimer
}

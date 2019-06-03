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
	// 如果是单独的timer服务，传false表示同步运行。传true表示异步运行，还可以在main里面调用其他的服务
	task.StartSchedule(tasks, false)
}

var demoTimer *task.TimeTask
func demoTimerTask() *task.TimeTask {
	if demoTimer == nil {
		demoTimer = &task.TimeTask{
			Name:"updateRedisDay",
			Run:service.UpdateDays, // 具体的业务逻辑
			Spec:"0 0 24 * * ?", // 触发时间
		}
	}
	return demoTimer
}

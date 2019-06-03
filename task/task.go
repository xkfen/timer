package task

import (
	"errors"
	"github.com/robfig/cron"
)

type TimeTask struct {
	Name string
	Run func()
	Spec string
}

func StartSchedule(tasks []*TimeTask, async bool) error {
	if len(tasks) == 0 {
		return errors.New("no task to do")
	}
	c := cron.New()
	for _, task := range tasks {
		c.AddFunc(task.Spec, task.Run)
	}
  if async {
    c.Run()
  }else {
    c.Start()
  }
	return nil
}

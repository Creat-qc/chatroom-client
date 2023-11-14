package chat

import (
	"github.com/liangdas/armyant/task"
	"io"
	"os"
)

type Manager struct {
	Writer io.Writer
}

func (this *Manager) writer() io.Writer {
	if this.Writer == nil {
		return os.Stdout
	}
	return this.Writer
}

func (this *Manager) Finish(task task.Task) {

}

func (this *Manager) CreateWork() task.Work {
	return NewWork(this)
}

func NewManger(t task.Task) task.WorkManager {
	this := new(Manager)
	return this
}

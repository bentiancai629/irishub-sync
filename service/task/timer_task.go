package task

// 重定义
type Command = func()

// task接口
type Task interface {
	GetCommand() Command
	GetCron() string
}

// task锁结构体
type LockTask struct {
	Spec     string
	cmd      Command
	withLock bool
}

// 构造task
func NewTask(spec string, cmd Command) Task {
	return &LockTask{
		Spec: spec,
		cmd:  cmd,
	}
}

// 通过环境变量构造task
func NewLockTaskFromEnv(spec string, cmd Command) Task {
	return NewTask(spec, cmd)
}

//实现task接口的lockTask对象
func (task *LockTask) GetCommand() Command {
	lockCmd := func() {
		task.cmd()
	}
	return lockCmd
}

//实现task接口的lockTask对象
func (task *LockTask) GetCron() string {
	return task.Spec
}

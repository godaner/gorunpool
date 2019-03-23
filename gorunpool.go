package gorunpool

type GoRunPool struct {
	InitConfig InitConfig
	tasks      chan Task
}

type InitConfig struct {
	Size int
}
type Callback func(output Params, err error)
type Process func(input Params) (output Params, err error)
type Task struct {
	ID       string
	Input    Params
	Process  Process
	Callback Callback
}
type Params map[string]interface{}

func NewPool(initConfig InitConfig) *GoRunPool {
	p := &GoRunPool{InitConfig: initConfig}
	p.startWorker()
	return p
}
func (this *GoRunPool) Run(task Task) {
	task.ID = NewUUID().String()
	this.tasks <- task
}

func (this *GoRunPool) startWorker() {
	this.tasks = make(chan Task, this.InitConfig.Size)
	for i := 1; i < this.InitConfig.Size; i++ {
		go func() {
			for task := range this.tasks {
				process := task.Process
				if nil != process {
					if task.Input == nil {
						task.Input = Params{}
					}
					task.Input["taskid"] = task.ID
					output, err := process(task.Input)
					callback := task.Callback
					if nil != callback {
						if output == nil {
							output = Params{}
						}
						output["taskid"] = task.ID
						callback(output, err)
					}
				}
			}

		}()
	}
}

package structures

type IContainer interface {
	AddTask(task Task)
	AddChild(container *Container)
}

type Container struct {
	Children []*Container
	Tasks    []Task
}

func (container *Container) AddTask(task Task) {
	container.Tasks = append(container.Tasks, task)
}

func (container *Container) AddChild(child *Container) {
	container.Children = append(container.Children, child)
}

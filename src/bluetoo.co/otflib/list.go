package otflib

type List struct {
	Name string
	Tasks []Task
}
func NewList(name string) *List {
	return &List{ Name: name }
}

func AddTask(task *Task, list *List) {
	list.Tasks = append(list.Tasks, *task)
}
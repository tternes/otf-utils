package otflib

type List struct {
	vendor Provider
	vendorListUid string

	Name string
	Tasks []Task
}
func NewList(name string) *List {
	return &List{ Name: name }
}

func AddTask(task *Task, list *List) {
	list.Tasks = append(list.Tasks, *task)
}
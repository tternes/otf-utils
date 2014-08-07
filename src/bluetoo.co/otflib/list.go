package otflib

type List struct {
	Name string
	Tasks []*Task
	
	// vendor fields
	VendorListId string
}
func NewList(name string) *List {
	return &List{ Name: name }
}

func AddTask(task *Task, list *List) {
	list.Tasks = append(list.Tasks, task)
}
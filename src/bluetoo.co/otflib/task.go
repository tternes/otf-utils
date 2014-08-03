package otflib

type Task struct {
	Name string
	IsCompleted bool
	IsDeleted bool
	IsStarred bool
}

func NewTask(name string) *Task {
	return &Task{ Name: name }
}

func (t *Task) SetCompleted(isComplete bool) {
	t.IsCompleted = isComplete
}

func (t *Task) SetDeleted(isDeleted bool) {
	t.IsDeleted = isDeleted
}

func (t *Task) SetStarred(isStarred bool) {
	t.IsStarred = isStarred
}
package otflib

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	
	list := NewList("list")
	if(len(list.Tasks) > 0) {
		t.Error("list should contain no tasks")
	}
	
	task := NewTask("first")
	AddTask(task, list)
	if(len(list.Tasks) != 1) {
		t.Error("list should contain one task")
	}
}

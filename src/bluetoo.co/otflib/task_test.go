package otflib

import (
	"testing"
)

func TestTaskCompleted(t *testing.T) {
	task := NewTask("task")
	
	if(task.IsCompleted) {
		t.Error("Task should default incomplete")
	}
	
	task.SetCompleted(true);
	if(!task.IsCompleted) {
		t.Error("Task should be completed after setting")
	}
}

func TestTaskStar(t *testing.T) {
	task := NewTask("task")
	if(task.IsStarred) {
		t.Error("Default star should be false")
	}
	
	task.SetStarred(true)
	if(!task.IsStarred) {
		t.Error("IsStarred should be true after setting")
	}
}
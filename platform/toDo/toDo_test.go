package toDo

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Task{})
	if len(feed.Tasks) != 1 {
		t.Errorf("Task was not added")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Task{})
	res := feed.GetAll()
	if len(res) != 1 {
		t.Errorf("Task was not added")
	}
}

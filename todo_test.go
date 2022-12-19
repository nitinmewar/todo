package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/nitinmewar/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskname := "get pizza"
	l.Add(taskname)

	if l[0].Task != taskname {
		t.Errorf("expected %s but got %s", taskname, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskname := "complete func"
	l.Add(taskname)

	if l[0].Task != taskname {
		t.Errorf("expected %s but got %s", taskname, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("new task should should mark completed")
	}

	l.Complete(1)
	if !l[0].Done {
		t.Errorf("task should be marked completed")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	taskname := "delete this "
	l.Add(taskname)

	if l[0].Task != taskname {
		t.Errorf("expected %s but got %s", taskname, l[0].Task)
	}

	l.Delete(1)
	if len(l) != 0 {
		t.Errorf("length should be 0 but got %d", len(l))
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskname := "testing"
	l1.Add(taskname)

	if l1[0].Task != taskname {
		t.Errorf("expected %s but got %s ", taskname, l1[0].Task)
	}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("error creating file")
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Errorf("error saving list to the file:  %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Errorf("error getting list from file:  %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("task %q should match  task %q", l1[0].Task, l2[0].Task)
	}
}

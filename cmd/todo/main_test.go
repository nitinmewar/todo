package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	filename = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("building tool")

	if runtime.GOOS == "windows" {
		binName += ".bin"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "can not build tools %s:  %s", binName, err)
		os.Exit(1)
	}
	fmt.Println("running tests")
	result := m.Run()

	fmt.Println("cleaning up")
	os.Remove(binName)
	os.Remove(filename)

	os.Exit(result)

}

func TestTodoCLI(t *testing.T) {
	task := "first testing task"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("addNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}
		expected := task + "\n"

		if expected != string(out) {
			t.Errorf("expecting %q but got %q", expected, string(out))
		}
	})
}

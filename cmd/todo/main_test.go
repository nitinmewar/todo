package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

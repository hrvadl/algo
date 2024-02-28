package tm

import (
	"os"
	"os/exec"
	"runtime"
)

func RunCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func Clear() {
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "linux":
		_ = RunCmd("clear")
	case "windows":
		_ = RunCmd("cmd", "/c", "cls")
	}
}

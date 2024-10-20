package utils

import (
	"os/exec"
	"runtime"
	"syscall"
)

func RunCmd(cmd *exec.Cmd) ([]byte, error) {
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	return cmd.CombinedOutput()
}

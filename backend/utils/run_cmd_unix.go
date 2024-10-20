//go:build !windows
// +build !windows

package utils

import (
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) ([]byte, error) {
	return cmd.CombinedOutput()
}

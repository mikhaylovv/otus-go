package hw7

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd - запускает программу с аргументами (cmd) c переопределнным окружением.
func RunCmd(cmd []string, env map[string]string) error {
	for key, val := range env {
		err := os.Setenv(key, val)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
		}
	}

	c := exec.Command(cmd[0], cmd[1:]...)
	envSlice := os.Environ()
	for key, val := range env {
		envSlice = append(envSlice, fmt.Sprintf("%s=%s", key, val))
	}
	c.Env = envSlice
	c.Stdout = os.Stdout
	return c.Run()
}

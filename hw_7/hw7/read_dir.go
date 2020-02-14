package hw7

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// ReadDir - сканирует указанный каталог и возвращает все переменные окружения, определенные в нем.
func ReadDir(dir string) (map[string]string, error) {
	d, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	files, err := d.Readdir(0)
	if err != nil {
		return nil, err
	}

	res := make(map[string]string, len(files))

	for _, file := range files {
		envName := file.Name()
		f, err := os.Open(dir + "/" + envName)
		if err != nil {
			continue
		}

		envVal, err := ioutil.ReadAll(f)
		if err != nil {
			continue
		}
		res[envName] = string(envVal)
	}

	return res, nil
}

// RunCmd - запускает программу с аргументами (cmd) c переопределнным окружением.
func RunCmd(cmd []string, env map[string]string) error {
	for key, val := range env {
		err := os.Setenv(key, val)
		if err != nil {
			os.Stderr.WriteString(err.Error())
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

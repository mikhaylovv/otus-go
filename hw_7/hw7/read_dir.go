package hw7

import (
	"bufio"
	"errors"
	"os"
	"path"
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
		if file.IsDir() {
			continue
		}

		envName := file.Name()
		f, err := os.Open(path.Join(dir, envName))
		if err != nil {
			continue
		}

		br := bufio.NewReader(f)

		envVal, isPrefix, err := br.ReadLine()
		if err != nil {
			continue
		}

		if isPrefix {
			return nil, errors.New("variable is too long for reading")
		}
		res[envName] = string(envVal)
	}

	return res, nil
}

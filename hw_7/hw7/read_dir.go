package hw7

import (
	"io/ioutil"
	"os"
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

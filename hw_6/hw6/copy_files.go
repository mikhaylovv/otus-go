package hw6

import (
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

/* Copy -
Реализовать утилиту копирования файлов Утилита должна принимать следующие аргументы:
 * файл источник (From)
 * файл копия (To)
 * Отступ в источнике (Offset), по умолчанию - 0
 * Количество копируемых байт (Limit),
по умолчанию - весь файл из From Выводить в консоль прогресс копирования в %,
например с помощью github.com/cheggaaa/pb Программа может НЕ обрабатывать файлы,
у которых не известна длинна (например /dev/urandom).

Завести в репозитории отдельный пакет (модуль) для этого ДЗ
Реализовать функцию вида Copy(from string, to string, limit int, offset int) error
Написать unit-тесты на функцию Copy
Реализовать функцию main, анализирующую параметры командной строки и вызывающую Copy
Проверить установку и работу утилиты руками
*/
func Copy(from string, to string, limit int64, offset int64) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	_, err = src.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	s, err := src.Stat()
	if err != nil {
		return err
	}
	fileSize := s.Size()

	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dst.Close()

	var barSize int64
	if size := fileSize - offset; size < limit || limit == 0 {
		barSize = size
	} else {
		barSize = limit
	}
	bar := pb.Full.Start64(barSize)
	defer bar.Finish()

	dstWriter := bar.NewProxyWriter(dst)

	if limit == 0 {
		_, err = io.Copy(dstWriter, src)
	} else {
		_, err = io.CopyN(dstWriter, src, limit)
	}

	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}


package hw3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTop10SimplePositive(t *testing.T) {
	as := assert.New(t)
	text := `one one
two two two
three three three three
four four four four four
five five five five five five
six six six six six six six
seven seven seven seven seven seven seven seven 
eight eight eight eight eight eight eight eight eight
nine nine nine nine nine nine nine nine nine nine
ten ten ten ten ten ten ten ten ten ten ten
here any other text for testing
`

	as.EqualValues(
		[10]string{"ten", "nine", "eight", "seven", "six", "five", "four", "three", "two", "one"},
		Top10(text))
}

func TestTop10HarderPositive(t *testing.T) {
	as := assert.New(t)
	as.EqualValues(
		[10]string{"и", "нога", "это", "слова", "разные", "с", "and", "dog", "go", "one"},
		Top10(`Частотный анализ
Цель: Напиcать функцию, принимающую на вход строку с текстом и возвращающую слайс с 10 самыми частовстречающимеся 
в тексте словами. Если есть более 10 самых частотых слов (например 15 разных слов встречаются ровно 133 раза, 
остальные < 100), можно вернуть любые 10 из самых частотных. Словоформы не учитываем. "нога", "ногу", 
"ноги" - это разные слова. Слово с большой и маленькой буквы можно считать за разные слова. "Нога" и 
"нога" - это разные слова. Знаки препиания можно считать "буквами" слова или отдельными словами. "-" (тире) - это 
отдельное слово. "нога," и "нога" - это разные слова. Пример: "cat and dog one dog two cats and one man". "dog", "one",
"and" - встречаются два раза, это топ-3. Задание со звездочкой (*): учитывать большие/маленьгие буквы и знаки
препинания. "Нога" и "нога" - это одинаковые слова, "нога," и "нога" - это одинаковые слова, "-" (тире) - это не слово.

Завести в репозитории отдельный пакет (модуль) для этого ДЗ
Реализовать функцию вид Top10(string) ([]string)
При необходимости выделить вспомогательные функции
Написать unit-тесты на функцию
Критерии оценки: Функция должна проходить все тесты
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать и проверить пакет с помощью go get / go test
Задание (*) НЕ влияет на баллы, оно дано просто для развития навыков. 
`))
}

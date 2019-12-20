package hw3

import (
	"strings"
	"unicode"
)

/*Top10 - Частотный анализ
Написать функцию, которая получает на вход текст и возвращает
10 самых часто встречающихся слов без учета словоформ
*/
func Top10(text string) []string {
	words := strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) })

	wordsMap := make(map[string]int, len(words))

	for _, word := range words {
		wordsMap[strings.ToLower(word)]++
	}

	res := []string{}
	for i := 0; i < 10 && len(wordsMap) > 0; i++ {
		// get top frequent word
		topFreqWord := ""
		it := 0
		for w, freq := range wordsMap {
			if freq > it {
				it = freq
				topFreqWord = w
			}
		}

		delete(wordsMap, topFreqWord)
		res = append(res, topFreqWord)
	}

	return res
}

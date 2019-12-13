package hw3

import (
	"strings"
	"unicode"
)

/*Top10 - Частотный анализ
Написать функцию, которая получает на вход текст и возвращает
10 самых часто встречающихся слов без учета словоформ
*/
func Top10(text string) [10]string {
	words := strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) })

	wordsMap := make(map[string]int, len(words))

	for _, word := range words {
		wordsMap[strings.ToLower(word)]++
	}

	var freq [10]string
	for i := 0; i < 10; i++ {
		topFreqWord := ""
		topFreq := 0
		for word, freq := range wordsMap {
			if freq > topFreq {
				topFreq = freq
				topFreqWord = word
			}
		}

		delete(wordsMap, topFreqWord)

		freq[i] = topFreqWord
	}

	return freq
}

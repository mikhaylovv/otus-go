package hw3

import (
	"strings"
	"unicode"
)

//getTopFreqWord - gets the most frequent word in dictionary
func getTopFreqWord(words map[string]int) string {
	word := ""
	it := 0
	for w, freq := range words {
		if freq > it {
			it = freq
			word = w
		}
	}

	return word
}

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

	freq := []string{}
	for i := 0; i < 10 && len(wordsMap) > 0; i++ {
		topFreqWord := getTopFreqWord(wordsMap)
		delete(wordsMap, topFreqWord)
		freq = append(freq, topFreqWord)
	}

	return freq
}

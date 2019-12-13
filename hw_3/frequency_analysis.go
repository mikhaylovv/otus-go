package hw3

import (
	"sort"
	"strings"
	"unicode"
)

//getTopFreqWord - gets the most frequent word in dictionary in lexicographical order
func getTopFreqWord (words map[string]int) string {
	word := ""
	it := 0
	for w, freq := range words {
		if freq > it {
			it = freq
			word = w
		}
	}

	top := []string{word}
	for word, freq := range words {
		if freq == it {
			top = append(top, word)
		}
	}

	sort.SliceStable(top, func (i, j int) bool { return top[i] < top[j] })

	return top[0]
}

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
		topFreqWord := getTopFreqWord(wordsMap)
		delete(wordsMap, topFreqWord)
		freq[i] = topFreqWord
	}

	return freq
}

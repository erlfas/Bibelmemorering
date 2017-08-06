package main

import (
	"strings"
	"unicode"
)

func htmlFilter(text string) string {
	return strings.Replace(text, "\n", "<br/>", -1)
}

func makeFourthTransformation(text string) string {
	words := getTokens(strings.Replace(text, "\n", "", -1))
	var wordsNew []string

	isWithinBoundaries := false
	for i, word := range words {
		if isBeginning(word, i, words) {
			if isWithinBoundaries {
				isWithinBoundaries = false
				wordsNew = append(wordsNew, "[...]")
			}

			if isWord, ending := isWord(word); isWord && len(strings.TrimSpace(word)) > 0 {
				wordsNew = append(wordsNew, string(word[0])+ending)
			} else {
				wordsNew = append(wordsNew, word)
			}
		} else {
			if isEnd(word) {
				wordsNew = append(wordsNew, "[...]")
				wordsNew = append(wordsNew, string(word[len(word)-1]))
				isWithinBoundaries = false
			} else {
				isWithinBoundaries = true
			}
		}
	}

	return strings.Join(wordsNew, " ")
}

func makeThirdTransformation(text string) string {
	words := getTokens(text)
	var wordsNew []string

	isWithinBoundaries := false
	for i, word := range words {
		if isAtBoundary(word, i, words) {
			if isWithinBoundaries {
				isWithinBoundaries = false
				wordsNew = append(wordsNew, "[...]")
			}

			if isWord, ending := isWord(word); isWord {
				wordsNew = append(wordsNew, string(word[0])+ending)
			} else {
				wordsNew = append(wordsNew, word)
			}
		} else {
			isWithinBoundaries = true
		}
	}

	return strings.Join(wordsNew, " ")
}

func makeSecondTransformation(text string) string {
	words := getTokens(text)
	var wordsNew []string

	for i, word := range words {
		if isAtBoundary(word, i, words) {
			if isWord, ending := isWord(word); isWord {
				wordsNew = append(wordsNew, string(word[0])+ending)
			} else {
				wordsNew = append(wordsNew, word)
			}
		} else {
			wordsNew = append(wordsNew, "_")
		}
	}

	return strings.Join(wordsNew, " ")
}

func makeFirstTransformation(text string) string {
	words := getTokens(text)
	var wordsNew []string
	for _, word := range words {
		wordTrimmed := strings.Trim(word, " ")
		if isWord, ending := isWord(wordTrimmed); isWord {
			wordsNew = append(wordsNew, string(wordTrimmed[0])+ending)
		} else {
			wordsNew = append(wordsNew, wordTrimmed)
		}
	}

	return strings.Join(wordsNew, " ")
}

func getTokens(text string) []string {
	words := strings.Split(text, " ")
	var words2 []string
	for _, w := range words {
		if len(w) > 0 {
			words2 = append(words2, w)
		}
	}
	return words2
}

func isBeginning(word string, i int, words []string) bool {
	if i == 0 {
		return true
	}

	nextLast := words[i-1]
	r := rune(nextLast[len(nextLast)-1])

	if r == '.' || r == '!' || r == ':' || unicode.IsDigit(r) || unicode.IsNumber(r) {
		return true
	}

	return false
}

func isAtBoundary(word string, i int, words []string) bool {
	if i == 0 {
		return true
	}

	if endsWithPunct(word) {
		return true
	}

	if endsWithPunct(words[i-1]) {
		return true
	}

	if isNumber(words[i-1]) {
		return true
	}

	return false
}

func isEnd(word string) bool {
	length := len(word)
	if length <= 0 {
		return false
	}

	r := rune(word[len(word)-1])

	return r == '.' || r == '!' || r == '\n' || r == ':'
}

func endsWithPunct(word string) bool {
	length := len(word)
	if length <= 0 {
		return false
	}

	r := rune(word[len(word)-1])

	return unicode.IsPunct(r) || r == '\n'
}

// isWord returns (isWord, endingPunctuationIfItIsAWord)
// endingPunctuationIfItIsAWord is "" if isWord = false
func isWord(word string) (bool, string) {
	length := len(word)

	if length <= 0 {
		return false, ""
	}

	isNumber := true
	for _, r := range word {
		if !unicode.IsNumber(r) && !unicode.IsDigit(r) {
			isNumber = false
		}
	}

	if isNumber {
		return false, ""
	}

	var runes []rune
	for i := len(word) - 1; i >= 0; i-- {
		r := rune(word[i])
		if erTegn(r) {
			runes = append(runes, r)
		} else {
			break
		}
	}

	ending := reverse(string(runes))

	return true, ending
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func erBokstav(r rune) bool {
	return unicode.IsLetter(r) || r == 'æ' || r == 'Æ' || r == 'ø' || r == 'Ø' || r == 'å' || r == 'Å'
}

func erTegn(r rune) bool {
	return unicode.IsPunct(r) || r == '\n'
}

func isNumber(word string) bool {
	for _, r := range word {
		if !unicode.IsNumber(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

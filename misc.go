package pm1

import "os"

func IsContain(r rune, runes []rune) bool {
	for _, rr := range runes {
		if rr == r {
			return true
		}
	}
	return false
}

func IsWithInRange(s string, ran []rune) bool {
	for _, r := range []rune(s) {
		if !IsContain(r, ran) {
			return false
		}
	}
	return true
}

func IsExistFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CalcMinimumWordLength(max, countOfWords int) int {
	// 区切りの数
	countOfSep := countOfWords - 1
	// 必要な各単語の合計の長さ
	wordsTotalLength := max - countOfSep
	// 各単語の最小の長さ
	wordLength := wordsTotalLength / countOfWords
	return wordLength
}

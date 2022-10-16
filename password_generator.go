package pm1

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

type PasswordGenerator struct {
	// alphabet english words
	// 文字の長さごとに配列を分ける
	words map[int][]string
}

func NewPasswordGenerator() *PasswordGenerator {
	return &PasswordGenerator{
		words: map[int][]string{},
	}
}

func (p *PasswordGenerator) Init() error {
	rand.Seed(time.Now().UnixNano())

	err := p.loadEnglishWordsFromTxt("assets/words_alpha.txt")
	if err != nil {
		return err
	}

	return nil
}

func (p *PasswordGenerator) loadEnglishWordsFromTxt(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	words := string(data)

	// split on newline
	for _, word := range strings.Split(words, "\n") {
		// もしその長さのキーがなかったら、配列を生成してあげる
		_, ok := p.words[len(word)]
		if !ok {
			p.words[len(word)] = []string{}
		}
		p.words[len(word)] = append(p.words[len(word)], word)
	}

	return nil
}

var (
	LowerAlphabets   = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	UpperAlphabets   = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	Numbers          = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	SupportedSymbols = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '-', '+', '|', '~', '?', '='}
)

func (p *PasswordGenerator) GeneratePassword(opt *PasswordOption) (string, error) {
	switch opt.Type {
	case Random:
		var runesUsing []rune
		runesUsing = append(runesUsing, LowerAlphabets...)
		runesUsing = append(runesUsing, UpperAlphabets...)
		if opt.UseNumber {
			runesUsing = append(runesUsing, Numbers...)
		}
		if opt.UseSymbol {
			runesUsing = append(runesUsing, opt.AllowedSymbols...)
		}

		password := ""
		for i := 0; i < opt.Length; i++ {
			password += string(runesUsing[rand.Intn(len(runesUsing))])
		}
		return password, nil
	}
	return "", nil
}

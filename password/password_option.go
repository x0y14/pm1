package password

type Option struct {
	Type Type

	// EasyToRemember
	AllowUpper   bool   // 大文字を許可するか
	MinLength    int    //　パスワードの最小の長さ
	MaxLength    int    // パスワードの最大の長さ
	CountOfWords int    // 単語数
	Separators   []rune // 単語の区切りに使用する文字

	// Random
	Length         int    // 生成するパスワードの長さ
	UseNumber      bool   // 数字を使用するか
	UseSymbol      bool   // 記号を使用するか
	AllowedSymbols []rune // 記号を使用する場合、使用してもよい記号たち
}

func isSupportedSymbol(symbol rune) bool {
	for _, supported := range SupportedSymbols {
		if supported == symbol {
			return true
		}
	}
	return false
}

func NewRandomOption(length int, useNumber, useSymbol bool, allowedSymbols []rune) *Option {
	var allowed []rune
	// 許可された記号から対応しているもののみを取り出す
	for _, allowedSymbol := range allowedSymbols {
		if isSupportedSymbol(allowedSymbol) {
			allowed = append(allowed, allowedSymbol)
		}
	}

	return &Option{
		Type:           Random,
		Length:         length,
		UseNumber:      useNumber,
		UseSymbol:      useSymbol,
		AllowedSymbols: allowed,
	}
}

func NewEasyToRememberOption(allowUpper bool, minLength, maxLength, countOfWords int, sep []rune) *Option {
	if minLength < 15 {
		minLength = 15
	}
	if maxLength < 15 {
		maxLength = 15
	}
	if maxLength < minLength {
		maxLength = minLength
	}

	if countOfWords < 3 {
		countOfWords = 3
	}

	if sep == nil {
		sep = SupportedSymbols
	}

	return &Option{
		Type:         EasyToRemember,
		AllowUpper:   allowUpper,
		MaxLength:    maxLength,
		CountOfWords: countOfWords,
		Separators:   sep,
	}
}

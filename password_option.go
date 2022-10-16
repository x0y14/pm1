package pm1

type PasswordOption struct {
	PassType PasswordType
	Length   int // 生成するパスワードの長さ

	// EasyToRemember
	Separator string // 単語の区切りに使用する文字列

	// Random
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

func NewRandomOption(length int, useNumber, useSymbol bool, allowedSymbols []rune) *PasswordOption {
	var allowed []rune
	// 許可された記号から対応しているもののみを取り出す
	for _, allowedSymbol := range allowedSymbols {
		if isSupportedSymbol(allowedSymbol) {
			allowed = append(allowed, allowedSymbol)
		}
	}

	return &PasswordOption{
		PassType:       Random,
		Length:         length,
		UseNumber:      useNumber,
		UseSymbol:      useSymbol,
		AllowedSymbols: allowed,
	}
}

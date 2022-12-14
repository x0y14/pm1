package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordGenerator_GeneratePassword_Random(t *testing.T) {
	gen := NewPasswordGenerator()
	err := gen.Init()
	if err != nil {
		t.Fatalf("failed to initialize generator: %v", err)
	}

	tests := []struct {
		name           string
		length         int
		useNumber      bool
		useSymbol      bool
		allowedSymbols []rune
	}{
		{
			"only u/l alpha",
			30,
			false,
			false,
			nil,
		},
		{
			"u/l alpha & number",
			40,
			true,
			false,
			nil,
		},
		{
			"u/l alpha & allowed symbols",
			20,
			false,
			true,
			[]rune{'!', '@'},
		},
		{
			"u/l alpha & all symbols",
			20,
			false,
			true,
			SupportedSymbols,
		},
		{
			"u/l alpha & number & allowed symbols",
			30,
			true,
			true,
			[]rune{'-', '+'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := NewRandomOption(tt.length, tt.useNumber, tt.useSymbol, tt.allowedSymbols)
			password, err := gen.GeneratePassword(opt)
			if err != nil {
				t.Fatalf("failed to generate password: %v", err)
			}

			// 生成されたパスワードの長さが期待するものと一致しているか
			assert.Equal(t, tt.length, len(password))

			// アルファベット大文字小文字のみ
			if !tt.useNumber && !tt.useSymbol {
				range_ := append(LowerAlphabets, UpperAlphabets...)
				if !IsWithInRange(password, range_) {
					t.Fatalf("password is not in range: %s", password)
				}
			}

			// アルファベット大文字小文字と数字のみ
			if tt.useNumber && !tt.useSymbol {
				range_ := append(LowerAlphabets, UpperAlphabets...)
				range_ = append(range_, Numbers...)
				if !IsWithInRange(password, range_) {
					t.Fatalf("password is not in range: %s", password)
				}
			}

			// アルファベット大文字小文字と数字と許可された記号のみ(全部)
			if tt.useNumber && tt.useSymbol {
				range_ := append(LowerAlphabets, UpperAlphabets...)
				range_ = append(range_, Numbers...)
				range_ = append(range_, tt.allowedSymbols...)
				if !IsWithInRange(password, range_) {
					t.Fatalf("password is not in range: %s", password)
				}
			}
		})
	}
}

func TestPasswordGenerator_GeneratePassword_EasyToRemember(t *testing.T) {
	gen := NewPasswordGenerator()
	err := gen.Init()
	if err != nil {
		t.Errorf("failed to initialize generator: %v", err)
	}

	tests := []struct {
		name         string
		allowUpper   bool
		minLength    int
		maxLength    int
		countOfWords int
		seps         []rune
	}{
		{
			"plus, 20 <= n <= 30",
			true,
			20,
			30,
			5,
			[]rune{'+'},
		},
		{
			"plus or hyphen, 30 <= n <= 50",
			true,
			30,
			50,
			7,
			[]rune{'+', '-'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set up
			gen := NewPasswordGenerator()
			err = gen.Init()
			if err != nil {
				t.Errorf("failed to initialize password generator: %v", err)
			}

			// option
			opt := NewEasyToRememberOption(tt.allowUpper, tt.minLength, tt.maxLength, tt.countOfWords, tt.seps)

			password, err := gen.GeneratePassword(opt)
			if err != nil {
				t.Errorf("failed to generate password: %v", err)
			}

			assert.True(t, tt.minLength <= len(password))
			assert.True(t, len(password) <= tt.maxLength)
			t.Logf("generated: %s(%d)", password, len(password))

		})
	}
}

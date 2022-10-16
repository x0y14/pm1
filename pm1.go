package pm1

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
)

// 参考: https://zenn.dev/yot0201/articles/6046138ec783d2

func GenIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func Pkcs7Pad(data []byte) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize)
	trail := bytes.Repeat([]byte{byte(length)}, length)
	return append(data, trail...)
}

func Enc(dataS, keyS string) (iv []byte, encrypted []byte, err error) {
	key, err := hex.DecodeString(keyS)
	if err != nil {
		return nil, nil, err
	}

	data, err := hex.DecodeString(dataS)
	if err != nil {
		return nil, nil, err
	}

	padded := Pkcs7Pad(data)
	encrypted = make([]byte, len(padded))

	iv, err = GenIV()
	if err != nil {
		return nil, nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, padded)

	return iv, encrypted, nil
}

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

func Encrypt(data []byte, key string) (encrypted, iv []byte, err error) {
	keyHex := make([]byte, hex.EncodedLen(len(key)))
	hex.Encode(keyHex, []byte(key))

	dataHex := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dataHex, data)

	paddedKey := Pkcs7Pad(keyHex)
	paddedData := Pkcs7Pad(dataHex)

	encrypted = make([]byte, len(paddedData))

	iv, err = GenIV()
	if err != nil {
		return nil, nil, err
	}

	block, err := aes.NewCipher(paddedKey)
	if err != nil {
		return nil, nil, err
	}

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, paddedData)

	return encrypted, iv, nil
}

func Pkcs7UnPad(data []byte) []byte {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	return data[:dataLength-padLength]
}

func Decrypt(encrypted []byte, key string, iv []byte) ([]byte, error) {
	keyHex := make([]byte, hex.EncodedLen(len(key)))
	hex.Encode(keyHex, []byte(key))
	paddedKey := Pkcs7Pad(keyHex)

	block, err := aes.NewCipher(paddedKey)
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(encrypted))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, encrypted)
	unPaddedHexData := Pkcs7UnPad(decrypted)

	data := make([]byte, hex.DecodedLen(len(unPaddedHexData)))
	_, err = hex.Decode(data, unPaddedHexData)
	if err != nil {
		return nil, err
	}
	return data, nil
}

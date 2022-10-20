package password

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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

func Encrypt(data []byte, hashedKey []byte) (encrypted, iv []byte, err error) {
	dataHex := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dataHex, data)

	paddedData := Pkcs7Pad(dataHex)

	encrypted = make([]byte, len(paddedData))

	iv, err = GenIV()
	if err != nil {
		return nil, nil, err
	}

	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, nil, err
	}

	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, paddedData)

	return encrypted, iv, nil
}

func Pkcs7UnPad(data []byte) ([]byte, error) {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	if dataLength-padLength <= 0 {
		return nil, fmt.Errorf("invalid data length: %d", dataLength-padLength)
	}
	return data[:dataLength-padLength], nil
}

func Decrypt(encrypted []byte, hashedKey []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(encrypted))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, encrypted)
	unPaddedHexData, err := Pkcs7UnPad(decrypted)
	if err != nil {
		return nil, err
	}

	data := make([]byte, hex.DecodedLen(len(unPaddedHexData)))
	_, err = hex.Decode(data, unPaddedHexData)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Sha256Hashing(data string) []byte {
	hashed := sha256.Sum256([]byte(data))
	return hashed[:]
}

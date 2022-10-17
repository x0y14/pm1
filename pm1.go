package pm1

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type EncryptedStorage struct {
	Body string `json:"body,omitempty"`
	Iv   string `json:"iv,omitempty"`
}

func Export(path string, encrypted, iv []byte) error {
	enc := EncryptedStorage{
		Body: hex.EncodeToString(encrypted),
		Iv:   hex.EncodeToString(iv),
	}

	encJson, err := json.Marshal(enc)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, encJson, 0700)
	if err != nil {
		return err
	}
	return nil
}

func Load(path string) (encrypted, iv []byte, err error) {
	if !IsExistFile(path) {
		return nil, nil, fmt.Errorf("file not found: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	var enc EncryptedStorage
	err = json.Unmarshal(data, &enc)
	if err != nil {
		return nil, nil, err
	}

	unHexedBody, err := hex.DecodeString(enc.Body)
	if err != nil {
		return nil, nil, err
	}

	unHexedIv, err := hex.DecodeString(enc.Iv)
	if err != nil {
		return nil, nil, err
	}

	return unHexedBody, unHexedIv, err
}

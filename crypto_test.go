package pm1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecrypt(t *testing.T) {
	data := []byte("hello, world")
	password := "passwd"
	encrypted, iv, err := Encrypt(data, Sha256Hashing(password))
	if err != nil {
		t.Fatalf("failed to encrypt data: %v", err)
	}

	result, err := Decrypt(encrypted, Sha256Hashing(password), iv)
	if err != nil {
		t.Fatalf("failed to decrypt data: %v", err)
	}

	assert.Equal(t, data, result)
}

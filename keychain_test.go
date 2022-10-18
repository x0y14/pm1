package pm1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testMasterPassword = "hello, pm1"
)

func TestGetHashedMasterPassword(t *testing.T) {
	err := SetHashedMasterPassword(
		Sha256Hashing(testMasterPassword),
		time.Now().Add(2*time.Minute))
	if err != nil {
		t.Errorf("failed to set mp: %v", err)
	}

	bytes, err := GetHashedMasterPassword()
	if err != nil {
		t.Errorf("failed to get mp: %v", err)
	}

	assert.Equal(t, Sha256Hashing(testMasterPassword), bytes)

	err = RemoveHashedMasterPassword()
	if err != nil {
		t.Errorf("failed to remove mp: %v", err)
	}
}

func TestGetHashedMasterPasswordExpired(t *testing.T) {
	expiredAt := time.Now().Add(-2 * time.Minute)
	err := SetHashedMasterPassword(
		Sha256Hashing(testMasterPassword), expiredAt)
	if err != nil {
		t.Errorf("failed to set mp: %v", err)
	}

	bytes, err := GetHashedMasterPassword()
	assert.Equal(t, fmt.Errorf("expired at: %s", expiredAt.Local().Format(time.RFC3339)), err)
	assert.Equal(t, []byte(nil), bytes)

	err = RemoveHashedMasterPassword()
	if err != nil {
		t.Errorf("failed to remove mp: %v", err)
	}
}

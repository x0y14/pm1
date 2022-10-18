package pm1

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	testStorage := NewStorage()
	testVault := NewVault("test")

	// for example.com
	// setup generator
	passwordGen := NewPasswordGenerator()
	err := passwordGen.Init()
	if err != nil {
		t.Errorf("failed to initialize password manager: %v", err)
	}
	// password option
	// in this case, the 40 length password will be generated using upper and lower alphabet and numbers.
	opt := NewRandomOption(40, true, false, nil)
	// generate password
	password, err := passwordGen.GeneratePassword(opt)
	if err != nil {
		t.Errorf("failed to generate password: %v", err)
	}

	conf := NewWebSiteConfidential("example", password, "https://example.com")
	testVault.Register(conf)

	// register vault onto testStorage
	testStorage.Register(testVault)

	// dump storage
	storageBytes, err := DumpStorage(testStorage)
	if err != nil {
		t.Errorf("failed to dump storage: %v", err)
	}

	// encryption
	masterPassword := "im master password"
	encryptedStorageBytes, iv, err := Encrypt(storageBytes, Sha256Hashing(masterPassword))
	if err != nil {
		t.Errorf("failed to encrypt storage: %v", err)
	}

	t.Logf("iv = %v\n", iv)
	t.Logf("encrypted = %v\n", encryptedStorageBytes)

	// write file
	temp, err := os.CreateTemp("", "")
	defer func(name string) {
		_ = os.Remove(name)
	}(temp.Name())

	if err != nil {
		t.Errorf("failed to create temp file: %v", err)
	}

	// enc, ivを出力する
	err = Export(temp.Name(), encryptedStorageBytes, iv)
	if err != nil {
		t.Errorf("failed to export: %v", err)
	}
	encryptedStorageBytes = nil
	iv = nil

	// enc, ivを読み取る
	encryptedStorageBytes, iv, err = Load(temp.Name())
	if err != nil {
		t.Errorf("failed to load: %v", err)
	}

	// decryption
	decryptedStorageBytes, err := Decrypt(encryptedStorageBytes, Sha256Hashing(masterPassword), iv)
	if err != nil {
		t.Errorf("faield to decrypt storage: %v", err)
	}

	// restore storage from bytes
	restoredTestStorage, err := LoadStorage(decryptedStorageBytes)
	if err != nil {
		t.Errorf("failed to restore storage: %v", err)
	}

	// check diff
	if diff := cmp.Diff(testStorage, restoredTestStorage); diff != "" {
		t.Errorf("restored storage mismatch (-want +got):\n%s", diff)
	}
}

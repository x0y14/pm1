package password

import (
	"fmt"
	"github.com/99designs/keyring"
	"time"
)

const (
	serviceName = "dev.x0y14.pm1"
)

func SetHashedMasterPassword(hashedMP []byte, endedAt time.Time) error {
	// キーチェーンを開く
	ring, err := keyring.Open(keyring.Config{
		ServiceName: serviceName,
	})
	if err != nil {
		return err
	}

	// ハッシュされたマスターパスワードを保存
	err = ring.Set(keyring.Item{
		Key:                         "hashed_master_password",
		Data:                        hashedMP,
		KeychainNotTrustApplication: false, // ようわからん
		KeychainNotSynchronizable:   true,  // 同期しない
	})
	if err != nil {
		return err
	}

	// 有効期間を保存
	err = ring.Set(keyring.Item{
		Key:                         "expired_at",
		Data:                        []byte(endedAt.Format(time.RFC3339)),
		KeychainNotTrustApplication: false, // ようわからん
		KeychainNotSynchronizable:   true,  // 同期しない
	})
	return err
}

func GetHashedMasterPassword() ([]byte, error) {
	// キーチェーンを開く
	ring, err := keyring.Open(keyring.Config{
		ServiceName: serviceName,
	})
	if err != nil {
		return nil, err
	}

	// 有効期限を取得
	expiredAtRaw, err := ring.Get("expired_at")
	if err != nil {
		return nil, err
	}
	expiredAt, err := time.Parse(time.RFC3339, string(expiredAtRaw.Data))
	if err != nil {
		return nil, err
	}
	// 有効期限切れ
	if !time.Now().Before(expiredAt) {
		return nil, fmt.Errorf("expired at: %s", expiredAt.Local().Format(time.RFC3339))
	}

	// ハッシュされたマスターパスワードを取得
	hashedMP, err := ring.Get("hashed_master_password")
	if err != nil {
		return nil, err
	}

	return hashedMP.Data, nil
}

func RemoveHashedMasterPassword() error {
	// キーチェーンを開く
	ring, err := keyring.Open(keyring.Config{
		ServiceName: serviceName,
	})
	if err != nil {
		return err
	}

	// ハッシュされたマスターパスワードを消す
	err = ring.Remove("hashed_master_password")
	if err != nil {
		return err
	}

	// 有効期限を消す
	err = ring.Remove("expired_at")
	return err
}

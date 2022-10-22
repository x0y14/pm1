package password

import (
	"encoding/json"
	"fmt"
)

type Storage struct {
	Vaults []*Vault `json:"vaults,omitempty"`
}

func NewStorage() *Storage {
	return &Storage{Vaults: []*Vault{}}
}

func (s *Storage) Register(v *Vault) {
	s.Vaults = append(s.Vaults, v)
}

func (s *Storage) FindVaultByName(name string) (*Vault, error) {
	for _, v := range s.Vaults {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, fmt.Errorf("vault not found: %v", name)
}

func DumpStorage(s *Storage) ([]byte, error) {
	return json.Marshal(s)
}

func LoadStorage(b []byte) (*Storage, error) {
	var storage Storage
	err := json.Unmarshal(b, &storage)
	if err != nil {
		return nil, err
	}
	return &storage, nil
}

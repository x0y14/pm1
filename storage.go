package pm1

import "encoding/json"

type Storage struct {
	Vaults []*Vault `json:"vaults,omitempty"`
}

func NewStorage() *Storage {
	return &Storage{Vaults: []*Vault{}}
}

func (s *Storage) Register(v *Vault) {
	s.Vaults = append(s.Vaults, v)
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

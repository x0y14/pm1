package pm1

import (
	"strings"
)

type Vault struct {
	Name string          `json:"name,omitempty"`
	Data []*Confidential `json:"data,omitempty"`
}

func NewVault(name string) *Vault {
	return &Vault{
		Name: name,
		Data: []*Confidential{},
	}
}

func (v *Vault) Register(conf *Confidential) {
	v.Data = append(v.Data, conf)
}

func (v *Vault) ListUpSameTarget(conf *Confidential) []*Confidential {
	var sameTargets []*Confidential
	for _, alreadyRegisteredConf := range v.Data {
		if alreadyRegisteredConf.Type == conf.Type {
			switch conf.Type {
			case WebSite:
				if alreadyRegisteredConf.Url == conf.Url {
					sameTargets = append(sameTargets, alreadyRegisteredConf)
				}
			case Application:
				if alreadyRegisteredConf.Identifier == conf.Identifier {
					sameTargets = append(sameTargets, alreadyRegisteredConf)
				}
			}
		}
	}
	return sameTargets
}

func (v *Vault) SearchWithName(name string) []*Confidential {
	var result []*Confidential
	for _, conf := range v.Data {
		if strings.Contains(conf.Name, name) {
			result = append(result, conf)
		}
	}
	return result
}

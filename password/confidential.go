package password

import "time"

type Confidential struct {
	Type      TargetType `json:"type,omitempty"`
	Name      string     `json:"name,omitempty"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	// WebSite
	Url string `json:"url,omitempty"`

	// Application
	Identifier string `json:"identifier,omitempty"`
}

func NewWebSiteConfidential(name, password, url string) *Confidential {
	return &Confidential{
		Type:      WebSite,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now().In(time.UTC),
		UpdatedAt: time.Now().In(time.UTC),
		Url:       url,
	}
}

func NewApplicationConfidential(name, password, identifier string) *Confidential {
	return &Confidential{
		Type:       Application,
		Name:       name,
		Password:   password,
		CreatedAt:  time.Now().In(time.UTC),
		UpdatedAt:  time.Now().In(time.UTC),
		Identifier: identifier,
	}
}

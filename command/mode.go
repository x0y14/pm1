package command

type Mode int

const (
	_ Mode = iota
	MHelp
	MVault
)

var modes = [...]string{
	MHelp:  "help",
	MVault: "vault",
}

func (m Mode) String() string {
	return modes[m]
}

func (m Mode) shortHelp() string {
	switch m {
	case MHelp:
		return shortHelp()
	case MVault:
		return vaultShortHelp()
	}
	return ""
}

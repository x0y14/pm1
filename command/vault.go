package command

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

func vaultHelp() string {
	var help string
	help += "about vault\n"
	help += "\tlist" + "\tdisplay a list of vaults\n"
	help += "\tcreate" + "\tcreate new vault\n"
	help += "\tdelete" + "\tdelete vault\n"

	help += "\n" + moreInfo("vault")
	return help
}

func vaultShortHelp() string {
	return "control vault"
}

func Vault() (*Command, error) {
	switch flag.Arg(1) {
	case "list":
		return &Command{
			Mode:     MVault,
			VaultOpt: &VaultOption{Subcommand: "list"},
		}, nil
	case "create":
	case "delete":
	}

	return nil, fmt.Errorf("unknown subcommand: \"%s\"\n%s", flag.Arg(1), moreInfoSubCmd("vault"))
}

type VaultOption struct {
	Subcommand string
}

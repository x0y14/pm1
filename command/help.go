package command

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

func Help() string {
	switch flag.Arg(1) {
	case "vault":
		return vaultHelp()
	}

	var help string
	help += "pm1\n"
	help += "the password manager\n\n"
	help += "available commands:\n"
	for i := 1; i < len(modes); i++ {
		help += fmt.Sprintf("\t%s\t%s\n", Mode(i).String(), Mode(i).shortHelp())
	}
	return help
}

func shortHelp() string {
	return "print help"
}

func moreInfo(cmd string) string {
	return fmt.Sprintf("use \"pm1 help %s\" for more information.", cmd)
}

func moreInfoSubCmd(cmd string) string {
	return fmt.Sprintf("use \"pm1 %s help <subcommand>\" for more information of subcommand.", cmd)
}

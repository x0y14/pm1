package command

import (
	flag "github.com/spf13/pflag"
)

func init() {
	// ヘルプ
	helpFlag := flag.NewFlagSet("help", flag.ExitOnError)

	// 保管庫に関して
	vaultFlag := flag.NewFlagSet("vault", flag.ExitOnError)
	helpFlag.AddFlagSet(vaultFlag)
	// 一覧
	vListFlag := flag.NewFlagSet("list", flag.ExitOnError)
	vaultFlag.AddFlagSet(vListFlag)
	// 作成
	vCreateFlag := flag.NewFlagSet("create", flag.ExitOnError)
	vaultFlag.AddFlagSet(vCreateFlag)
}

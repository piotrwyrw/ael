package commands

import (
	"ael/internal/config"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func CmdInitialize(ctx context.Context, cmd *cli.Command) error {
	if config.DoesConfigExist() {
		return cli.Exit("An AEL configuration already exists in this directory. Please remove it ("+config.AelConfigFile+") before proceeding", 1)
	}

	err := config.StoreEmptyConfiguration()
	if err != nil {
		return err
	}

	fmt.Println("Initialized empty AEL configuration in current directory")

	return nil
}

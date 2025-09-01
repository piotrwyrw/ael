package cli

import (
	"ael/internal/cli/commands"
	"ael/internal/config"
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func InitializeCLI(config *config.AelConfig) error {
	ctx := commands.AelContext{Config: config}

	initCommand := cli.Command{
		Name:   "init",
		Usage:  "Initialize a new AEL Configuration",
		Action: commands.CmdInitialize,
	}

	nodeCommand := cli.Command{
		Name:  "node",
		Usage: "List or Manage Aether Nodes",
		Commands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "List All Registered Aether Nodes",
				Action: ctx.CmdListNodes,
			},
			{
				Name:  "add",
				Usage: "Add a New Aether Node",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the New Aether Node Entry",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "address",
						Aliases:  []string{"a"},
						Usage:    "Address of the Aether Node",
						Required: true,
					},
				},
				Action: ctx.CmdAddNode,
			},
			{
				Name:  "remove",
				Usage: "Remove an Existing Aether Node",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the Aether Node to Remove",
						Required: true,
					},
				},
				Action: ctx.CmdRemoveNode,
			},
		},
	}

	masterCommand := &cli.Command{
		Name:  "ael",
		Usage: "Aether Execution Layer",
		Commands: []*cli.Command{
			&initCommand,
			&nodeCommand,
		},
	}

	err := masterCommand.Run(context.Background(), os.Args)
	return err
}

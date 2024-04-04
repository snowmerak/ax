package main

import (
	"errors"
	"github.com/urfave/cli/v2"
)

func dockerRun(ctx *cli.Context) error {
	args := ctx.Args().Slice()

	if len(args) == 0 {
		return errors.New("no container image specified")
	}

	cfg, err := loadContainerConfig()
	if err != nil {
		return err
	}

	if err := run(cfg.Engine, append([]string{"run", "-v", ".:/workspace", "-w", "/workspace", cfg.Aliases[args[0]]}, args[1:]...)...); err != nil {
		return err
	}

	return nil
}

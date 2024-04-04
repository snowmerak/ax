package main

import "github.com/urfave/cli/v2"

func dockerRunBuf(ctx *cli.Context) error {
	args := ctx.Args().Slice()

	cfg, err := loadContainerConfig()
	if err != nil {
		return err
	}

	if err := run(cfg.Engine, append([]string{"run", "-v", ".:/workspace", "-w", "/workspace", "bufbuild/buf"}, args...)...); err != nil {
		return err
	}

	return nil
}

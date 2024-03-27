package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func commit(ctx *cli.Context) error {
	message := ctx.Args().First()

	if message == "" {
		return fmt.Errorf("commit message is required")
	}

	if err := run("git", "add", "."); err != nil {
		return fmt.Errorf("failed to add changes: %w", err)
	}

	if err := run("git", "commit", "-m", message); err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}

	return nil
}

func gitPush() error {
	if err := run("git", "push"); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}

	return nil
}

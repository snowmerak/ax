package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func push(ctx *cli.Context) error {
	branchName, err := getBranchName()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	sp := strings.Split(branchName, "/")

	switch len(sp) {
	case 4:
		if err := gitPush(); err != nil {
			return fmt.Errorf("failed to push changes: %w", err)
		}

		if err := run("git", "switch", fmt.Sprintf("%s/%s", sp[0], sp[1])); err != nil {
			return fmt.Errorf("failed to switch to feature branch: %w", err)
		}
	case 2:
		if err := gitPush(); err != nil {
			return fmt.Errorf("failed to push changes: %w", err)
		}

		if err := run("git", "switch", "unstable"); err != nil {
			return fmt.Errorf("failed to switch to unstable branch: %w", err)
		}
	case 1:
		switch sp[0] {
		case "unstable":
			if err := gitPush(); err != nil {
				return fmt.Errorf("failed to push changes: %w", err)
			}

			if err := run("git", "switch", "stable"); err != nil {
				return fmt.Errorf("failed to switch to stable branch: %w", err)
			}
		case "stable":
			if err := gitPush(); err != nil {
				return fmt.Errorf("failed to push changes: %w", err)
			}

			if err := run("git", "switch", "prod"); err != nil {
				return fmt.Errorf("failed to switch to prod branch: %w", err)
			}
		case "prod":
			return fmt.Errorf("you are already on the prod branch")
		}
	default:
		return fmt.Errorf("invalid branch name")
	}

	return nil
}

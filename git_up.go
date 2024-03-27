package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func up(ctx *cli.Context) error {
	branchName, err := getBranchName()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	sp := strings.Split(branchName, "/")

	switch len(sp) {
	case 4:
		if err := run("git", "switch", fmt.Sprintf("%s/%s", sp[0], sp[1])); err != nil {
			return fmt.Errorf("failed to switch to feature branch: %w", err)
		}
	case 2:
		if err := run("git", "switch", "develop"); err != nil {
			return fmt.Errorf("failed to switch to develop branch: %w", err)
		}
	case 1:
		switch sp[0] {
		case "develop":
			if err := run("git", "switch", "staging"); err != nil {
				return fmt.Errorf("failed to switch to staging branch: %w", err)
			}
		case "staging":
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

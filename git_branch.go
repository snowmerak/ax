package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func switchFeature(ctx *cli.Context) error {
	featureName := ctx.Args().First()

	if featureName == "" {
		return fmt.Errorf("feature name is required")
	}

	featureName = fmt.Sprintf("feature/%s", featureName)

	if err := run("git", "switch", "develop"); err != nil {
		return fmt.Errorf("failed to switch to develop branch: %w", err)
	}

	if err := run("git", "pull"); err != nil {
		return fmt.Errorf("failed to pull the repository: %w", err)
	}

	if err := run("git", "switch", "-C", featureName); err != nil {
		return fmt.Errorf("failed to create feature branch: %w", err)
	}

	if err := run("git", "push", "-u", "origin", featureName); err != nil {
		return fmt.Errorf("failed to push to remote repository: %w", err)
	}

	return nil
}

func switchProposal(ctx *cli.Context) error {
	proposalName := ctx.Args().First()

	if proposalName == "" {
		return fmt.Errorf("proposal name is required")
	}

	name, err := getBranchName()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	if !(strings.HasPrefix(name, "feature/") || strings.HasPrefix(name, "bugfix/")) {
		return fmt.Errorf("you must be on a feature or bugfix branch")
	}

	proposalName = fmt.Sprintf("%s/proposal/%s", name, proposalName)

	if err := run("git", "switch", "-C", proposalName); err != nil {
		return fmt.Errorf("failed to create proposal branch: %w", err)
	}

	if err := run("git", "push", "-u", "origin", proposalName); err != nil {
		return fmt.Errorf("failed to push to remote repository: %w", err)
	}

	return nil
}

func switchBugfix(ctx *cli.Context) error {
	bugfixName := ctx.Args().First()

	if bugfixName == "" {
		return fmt.Errorf("bugfix name is required")
	}

	bugfixName = fmt.Sprintf("bugfix/%s", bugfixName)

	if err := run("git", "switch", "develop"); err != nil {
		return fmt.Errorf("failed to switch to develop branch: %w", err)
	}

	if err := run("git", "pull"); err != nil {
		return fmt.Errorf("failed to pull the repository: %w", err)
	}

	if err := run("git", "switch", "-C", bugfixName); err != nil {
		return fmt.Errorf("failed to create bugfix branch: %w", err)
	}

	if err := run("git", "push", "-u", "origin", bugfixName); err != nil {
		return fmt.Errorf("failed to push to remote repository: %w", err)
	}

	return nil
}

func switchHotfix(ctx *cli.Context) error {
	hotfixName := ctx.Args().First()

	if hotfixName == "" {
		return fmt.Errorf("hotfix name is required")
	}

	hotfixName = fmt.Sprintf("hotfix/%s", hotfixName)

	if err := run("git", "switch", "prod"); err != nil {
		return fmt.Errorf("failed to switch to prod branch: %w", err)
	}

	if err := run("git", "pull"); err != nil {
		return fmt.Errorf("failed to pull the repository: %w", err)
	}

	if err := run("git", "switch", "-C", hotfixName); err != nil {
		return fmt.Errorf("failed to create hotfix branch: %w", err)
	}

	if err := run("git", "push", "-u", "origin", hotfixName); err != nil {
		return fmt.Errorf("failed to push to remote repository: %w", err)
	}

	return nil
}

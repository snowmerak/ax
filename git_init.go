package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func initRepo(ctx *cli.Context) error {
	remoteRepogitory := ctx.Args().First()

	if remoteRepogitory == "" {
		return fmt.Errorf("error: remote repository is required")
	}

	if err := run("git", "init"); err != nil {
		return fmt.Errorf("error initializing git repository: %w", err)
	}

	if err := run("git", "switch", "-C", "release"); err != nil {
		return fmt.Errorf("error creating release branch: %w", err)
	}

	if err := os.WriteFile("README.md", []byte("# CHANGE ME!"), 0644); err != nil {
		return fmt.Errorf("error creating README.md: %w", err)
	}

	if err := run("git", "add", "README.md"); err != nil {
		return fmt.Errorf("error adding README.md: %w", err)
	}

	if err := run("git", "commit", "-m", "Initial commit"); err != nil {
		return fmt.Errorf("error committing README.md: %w", err)
	}

	if err := run("git", "remote", "add", "origin", remoteRepogitory); err != nil {
		return fmt.Errorf("error adding remote repository: %w", err)
	}

	if err := run("git", "push", "-u", "origin", "prod"); err != nil {
		return fmt.Errorf("error pushing to remote repository: %w", err)
	}

	if err := run("git", "switch", "-C", "stable"); err != nil {
		return fmt.Errorf("error creating stable branch: %w", err)
	}

	if err := run("git", "push", "-u", "origin", "stable"); err != nil {
		return fmt.Errorf("error pushing to remote repository: %w", err)
	}

	if err := run("git", "switch", "-C", "unstable"); err != nil {
		return fmt.Errorf("error creating unstable branch: %w", err)
	}

	if err := run("git", "push", "-u", "origin", "unstable"); err != nil {
		return fmt.Errorf("error pushing to remote repository: %w", err)
	}

	return nil
}

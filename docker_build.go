package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func buildDockerImage(ctx *cli.Context) error {
	cfg, err := loadContainerConfig()
	if err != nil {
		return fmt.Errorf("error loading container config: %w", err)
	}

	options := make([]string, len(cfg.Registries))
	for i, r := range cfg.Registries {
		options[i] = fmt.Sprintf("%d: %s.Dockerfile -t %s:%s", i+1, r.AppName, r.Registry, r.Tag)
	}

	selected := []string{}
	if err := survey.AskOne(&survey.MultiSelect{
		Message: "Select the registty and tag",
		Options: options,
	}, &selected, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("error asking questions: %w", err)
	}

	selectedIndexes := make([]int, len(selected))
	fmt.Println("Selected registries:")
	for i, s := range selected {
		fmt.Println(s)
		sp := strings.SplitN(s, ":", 2)
		idx, err := strconv.Atoi(sp[0])
		if err != nil {
			return fmt.Errorf("error parsing index: %w", err)
		}
		selectedIndexes[i] = idx - 1
	}

	major := 0
	if ctx.Bool("major") {
		major = 1
	}
	minor := 0
	if ctx.Bool("minor") {
		minor = 1
	}
	patch := 0
	if ctx.Bool("patch") {
		patch = 1
	}

	for _, idx := range selectedIndexes {
		file := fmt.Sprintf("%s.Dockerfile", cfg.Registries[idx].AppName)
		tag := fmt.Sprintf("%s:%s", filepath.Join(cfg.Registries[idx].Registry, pascal2Snake(cfg.Registries[idx].AppName)), cfg.Registries[idx].Tag)
		if err := buildImage(file, tag); err != nil {
			return fmt.Errorf("error building image: %w", err)
		}
		if cfg.Registries[idx].TagAutoIncrement {
			err := error(nil)
			cfg.Registries[idx].Tag, err = incrementTag(cfg.Registries[idx].Tag, major, minor, patch)
			if err != nil {
				return fmt.Errorf("error incrementing tag: %w", err)
			}

			if err := saveContainerConfig(cfg); err != nil {
				return fmt.Errorf("error saving container config: %w", err)
			}

			if ctx.Bool("push") {
				if err := run("docker", "push", tag); err != nil {
					return fmt.Errorf("error pushing image: %w", err)
				}
			}
		}
	}

	return nil
}

func buildImage(file, tag string) error {
	if err := run("docker", "build", "-t", tag, "-f", file, "."); err != nil {
		return fmt.Errorf("error building image: %w", err)
	}
	return nil
}

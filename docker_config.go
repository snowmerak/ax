package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/dlclark/regexp2"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

const containerConfigFile = "container.yaml"

type ContainerRegistry struct {
	Tag              string `survey:"tag" yaml:"tag"`
	Registry         string `survey:"registry" yaml:"registry"`
	AppName          string `survey:"app_name" yaml:"app_name"`
	TagAutoIncrement bool   `survey:"tag_auto_increment" yaml:"tag_auto_increment"`
}

type ContainerConfig struct {
	Engine     string               `survey:"engine" yaml:"engine"`
	Aliases    map[string]string    `yaml:"aliases"`
	Registries []*ContainerRegistry `survey:"registries" yaml:"registries"`
}

func initContainerConfig(ctx *cli.Context) error {
	qs := []*survey.Question{
		{
			Name: "engine",
			Prompt: &survey.Select{
				Message: "Select the container engine",
				Options: []string{"docker", "podman"},
			},
			Validate: survey.Required,
		},
	}

	cc := ContainerConfig{}
	if err := survey.Ask(qs, &cc); err != nil {
		return fmt.Errorf("error asking questions: %w", err)
	}

	{
		ok := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "Do you want to add container registries?",
		}, &ok); err != nil {
			return fmt.Errorf("error asking questions: %w", err)
		}

		if ok {
			registries, err := inputRegistries()
			if err != nil {
				return fmt.Errorf("error asking questions: %w", err)
			}

			cc.Registries = registries
		}
	}

	cc.Aliases = map[string]string{
		"buf":  "bufbuild/buf",
		"sqlc": "sqlc/sqlc",
	}

	if err := saveContainerConfig(&cc); err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}

	return nil
}

func inputRegistries() ([]*ContainerRegistry, error) {
	registries := []*ContainerRegistry{}
	qs := []*survey.Question{
		{
			Name: "app_name",
			Prompt: &survey.Input{
				Message: "Enter the name of the app",
			},
			Validate: survey.Required,
		},
		{
			Name: "registry",
			Prompt: &survey.Input{
				Message: "Enter the name of the registry",
			},
			Validate: survey.Required,
		},
		{
			Name: "tag",
			Prompt: &survey.Input{
				Message: "Enter the tag",
			},
			Validate: survey.Required,
		},
		{
			Name: "tag_auto_increment",
			Prompt: &survey.Confirm{
				Message: "Do you want to auto-increment the tag?",
			},
			Validate: survey.Required,
		},
	}

	for {
		cr := ContainerRegistry{}
		if err := survey.Ask(qs, &cr); err != nil {
			return nil, fmt.Errorf("error asking questions: %w", err)
		}

		if err := validateContainerTag(cr.Tag); err != nil {
			return nil, fmt.Errorf("error validating tag: %w", err)
		}

		registries = append(registries, &cr)

		ok := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "Do you want to add another registry?",
		}, &ok); err != nil {
			return nil, fmt.Errorf("error asking questions: %w", err)
		}

		if !ok {
			break
		}
	}

	return registries, nil
}

func loadContainerConfig() (*ContainerConfig, error) {
	f, err := os.ReadFile(containerConfigFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	cc := ContainerConfig{}
	if err := yaml.Unmarshal(f, &cc); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &cc, nil
}

func saveContainerConfig(cc *ContainerConfig) error {
	cfg, err := yaml.Marshal(cc)
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}

	if err := os.WriteFile(containerConfigFile, cfg, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func validateContainerTag(tag string) error {
	if tag == "" {
		return fmt.Errorf("tag cannot be empty")
	}

	switch tag {
	case "dev", "staging":
		return nil
	}

	regex, err := regexp2.Compile(`^\d+\.\d+\.\d+$`, regexp2.DefaultUnmarshalOptions)
	if err != nil {
		return fmt.Errorf("error compiling regex: %w", err)
	}

	if ok, err := regex.MatchString(tag); err != nil {
		return fmt.Errorf("error matching regex: %w", err)
	} else if !ok {
		return fmt.Errorf("tag must be in the format x.y.z")
	}

	return nil
}

func incrementTag(tag string, majorVariance int, minorVariance int, patchVariance int) (string, error) {
	switch tag {
	case "dev", "staging":
		return tag, nil
	}

	regex, err := regexp2.Compile(`^(\d+)\.(\d+)\.(\d+)$`, regexp2.DefaultUnmarshalOptions)
	if err != nil {
		return "", fmt.Errorf("error compiling regex: %w", err)
	}

	m, err := regex.FindStringMatch(tag)
	if err != nil {
		return "", fmt.Errorf("error finding match: %w", err)
	}

	major, err := strconv.Atoi(m.GroupByNumber(1).String())
	if err != nil {
		return "", fmt.Errorf("error parsing major: %w", err)
	}

	minor, err := strconv.Atoi(m.GroupByNumber(2).String())
	if err != nil {
		return "", fmt.Errorf("error parsing minor: %w", err)
	}

	patch, err := strconv.Atoi(m.GroupByNumber(3).String())
	if err != nil {
		return "", fmt.Errorf("error parsing patch: %w", err)
	}

	major += majorVariance
	minor += minorVariance
	patch += patchVariance

	return fmt.Sprintf("%d.%d.%d", major, minor, patch), nil
}

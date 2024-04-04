package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "AX",
		Usage: "AX is a simple CLI tool for managing your project",
		Commands: []*cli.Command{
			{
				Name:    "git",
				Aliases: []string{"g"},
				Subcommands: []*cli.Command{
					{
						Name:      "init",
						Aliases:   []string{"i"},
						Usage:     "Initialize a new git repository",
						UsageText: "ax init <remote-repository>",
						Action:    initRepo,
					}, {
						Name:      "push",
						Aliases:   []string{"u"},
						Usage:     "Push and switch back to the previous branch",
						UsageText: "ax push",
						Action:    push,
					}, {
						Name:      "commit",
						Aliases:   []string{"c"},
						Usage:     "Commit changes",
						UsageText: "ax commit <message>",
						Action:    commit,
					}, {
						Name:      "feature",
						Aliases:   []string{"f"},
						Usage:     "Create a new feature branch",
						UsageText: "ax feature <feature-name>",
						Action:    switchFeature,
					}, {
						Name:      "bugfix",
						Aliases:   []string{"b"},
						Usage:     "Create a new bugfix branch",
						UsageText: "ax bugfix <bugfix-name>",
						Action:    switchBugfix,
					}, {
						Name:      "hotfix",
						Aliases:   []string{"x"},
						Usage:     "Create a new hotfix branch",
						UsageText: "ax hotfix <hotfix-name>",
						Action:    switchHotfix,
					}, {
						Name:      "proposal",
						Aliases:   []string{"p"},
						Usage:     "Create a new proposal branch",
						UsageText: "ax proposal <proposal-name>",
						Action:    switchProposal,
					}, {
						Name:      "develop",
						Aliases:   []string{"d"},
						Usage:     "Switch to the develop branch",
						UsageText: "ax develop",
						Action:    switchDevelop,
					}, {
						Name:      "staging",
						Aliases:   []string{"s"},
						Usage:     "Switch to the staging branch",
						UsageText: "ax staging",
						Action:    switchStaging,
					},
				},
			},
			{
				Name:    "container",
				Aliases: []string{"c"},
				Subcommands: []*cli.Command{
					{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "Configure container",
						Subcommands: []*cli.Command{
							{
								Name:      "init",
								Aliases:   []string{"i"},
								Usage:     "Initialize container config",
								UsageText: "ax container config init",
								Action:    initContainerConfig,
							},
						},
					},
					{
						Name:    "image",
						Aliases: []string{"i"},
						Subcommands: []*cli.Command{
							{
								Name:      "init",
								Aliases:   []string{"i"},
								Usage:     "Initialize docker image",
								UsageText: "ax container image init <language-option> <name>",
								Action:    initDockerImage,
								Flags: []cli.Flag{
									&cli.BoolFlag{
										Name:     "go",
										Aliases:  []string{"g"},
										Usage:    "Initialize a Go project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "node",
										Aliases:  []string{"n"},
										Usage:    "Initialize a Node.js project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "python",
										Aliases:  []string{"p"},
										Usage:    "Initialize a Python project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "jdk",
										Aliases:  []string{"j"},
										Usage:    "Initialize a JDK project image",
										Value:    false,
										Category: "language",
									},
								},
							},
							{
								Name:      "build",
								Aliases:   []string{"b"},
								Usage:     "Build docker image",
								UsageText: "ax container image build <name>",
								Action:    buildDockerImage,
								Flags: []cli.Flag{
									&cli.BoolFlag{
										Name:    "push",
										Aliases: []string{"p"},
										Usage:   "Push the image to the registry",
									},
								},
							},
						},
					},
					{
						Name:    "devcontainer",
						Aliases: []string{"d"},
						Subcommands: []*cli.Command{
							{
								Name:      "init",
								Aliases:   []string{"i"},
								Usage:     "Initialize devcontainer",
								UsageText: "ax container devcontainer init <language-option> <name>",
								Action:    initDevContainer,
								Flags: []cli.Flag{
									&cli.BoolFlag{
										Name:     "go",
										Aliases:  []string{"g"},
										Usage:    "Initialize a Go project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "node",
										Aliases:  []string{"n"},
										Usage:    "Initialize a Node.js project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "python",
										Aliases:  []string{"p"},
										Usage:    "Initialize a Python project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "jdk",
										Aliases:  []string{"j"},
										Usage:    "Initialize a JDK project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "cpp",
										Aliases:  []string{"c"},
										Usage:    "Initialize a C++ project image",
										Value:    false,
										Category: "language",
									},
									&cli.BoolFlag{
										Name:     "rust",
										Aliases:  []string{"r"},
										Usage:    "Initialize a Rust project image",
										Value:    false,
										Category: "language",
									},
								},
							},
						},
					},
					{
						Name:    "run",
						Aliases: []string{"r"},
						Usage:   "Run container",
						Subcommands: []*cli.Command{
							{
								Name:      "buf",
								Aliases:   []string{"b"},
								Usage:     "Run buf container",
								UsageText: "ax container run buf <args>",
								Action:    dockerRunBuf,
							},
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

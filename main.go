package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "AX",
		Usage:     "AX is a simple CLI tool for managing your project",
		UsageText: "ax command [command options] [arguments...]",
		Commands: []*cli.Command{
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
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

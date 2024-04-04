package main

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

func initDevContainer(ctx *cli.Context) error {
	name := ctx.Args().First()

	dir := filepath.Join(".devcontainer", name)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating .devcontainer directory: %w", err)
	}

	dockerFileBuilder := bytes.NewBuffer(nil)
	dockerFileBuilder.WriteString("FROM ")
	switch {
	case ctx.Bool("go"):
		writeGoDevContainer(dockerFileBuilder)
	case ctx.Bool("node"):
		writeNodeDevContainer(dockerFileBuilder)
	case ctx.Bool("python"):
		writePythonDevContainer(dockerFileBuilder)
	}

	dockerFilePath := filepath.Join(dir, "Dockerfile")
	if err := os.WriteFile(dockerFilePath, dockerFileBuilder.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing Dockerfile: %w", err)
	}

	devcontainerFilePath := filepath.Join(dir, "devcontainer.json")
	devcontainerFileContent := bytes.NewBuffer(nil)
	devcontainerFileContent.WriteString("{\n")
	devcontainerFileContent.WriteString("  \"name\": \"")
	devcontainerFileContent.WriteString(name)
	devcontainerFileContent.WriteString("\",\n")
	devcontainerFileContent.WriteString("  \"dockerFile\": \"Dockerfile\"\n")
	devcontainerFileContent.WriteString("}\n")
	if err := os.WriteFile(devcontainerFilePath, devcontainerFileContent.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing devcontainer.json: %w", err)
	}

	return nil
}

func writeGoDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString("golang:1.22-bookworm\n")
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writeNodeDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString("node:20-bookworm\n")
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writePythonDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString("python:3.12-bookworm\n")
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}
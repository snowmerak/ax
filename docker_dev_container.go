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
	case ctx.Bool("jdk"):
		writeJdkDevContainer(dockerFileBuilder)
	case ctx.Bool("cpp"):
		writeCppDevContainer(dockerFileBuilder)
	case ctx.Bool("rust"):
		writeRustDevContainer(dockerFileBuilder)
	default:
		return fmt.Errorf("no language specified")
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
	devcontainerFileContent.WriteString("  \"dockerComposeFile\": [\n")
	devcontainerFileContent.WriteString("    \"docker-compose.yml\"\n")
	devcontainerFileContent.WriteString("  ],\n")
	devcontainerFileContent.WriteString("  \"service\": \"")
	devcontainerFileContent.WriteString(name)
	devcontainerFileContent.WriteString("\",\n")
	devcontainerFileContent.WriteString("  \"workspaceFolder\": \"/workspace\",\n")
	devcontainerFileContent.WriteString("  \"shutdownAction\": \"stopCompose\"\n")
	devcontainerFileContent.WriteString("}\n")
	if err := os.WriteFile(devcontainerFilePath, devcontainerFileContent.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing devcontainer.json: %w", err)
	}

	devComposeFilePath := filepath.Join(dir, "docker-compose.yml")
	devComposeFileContent := bytes.NewBuffer(nil)
	devComposeFileContent.WriteString("version: '3'\n")
	devComposeFileContent.WriteString("services:\n")
	devComposeFileContent.WriteString("  ")
	devComposeFileContent.WriteString(name)
	devComposeFileContent.WriteString(":\n")
	devComposeFileContent.WriteString("    build:\n")
	devComposeFileContent.WriteString("      context: .\n")
	devComposeFileContent.WriteString("      dockerfile: Dockerfile\n")
	devComposeFileContent.WriteString("    volumes:\n")
	devComposeFileContent.WriteString("      - ../../.:/workspace:cached\n")
	if err := os.WriteFile(devComposeFilePath, devComposeFileContent.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing docker-compose.yml: %w", err)
	}

	return nil
}

const UbuntuBaseImage = "ubuntu:24.04"
const DebianAlias = "bookworm"
const JdkVersion = "21"
const PythonVersion = "3.12"
const NodeVersion = "20"
const GolangVersion = "1.22"
const RustVersion = "1.77"

func writeGoDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("golang:%s-%s\n\n", GolangVersion, DebianAlias))
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writeNodeDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("node:%s-%s\n\n", NodeVersion, DebianAlias))
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writePythonDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("python:%s-%s\n\n", PythonVersion, DebianAlias))
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writeJdkDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("mcr.microsoft.com/openjdk/jdk:%s-ubuntu\n\n", JdkVersion))
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writeCppDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString(UbuntuBaseImage)
	buffer.WriteString("\n\n")
	buffer.WriteString("RUN apt update && apt install -y build-essential autoconf autoconf-archive binutils ninja-build curl file gcc g++ git libtool make musl-dev tar unzip zip wget pkg-config\n\n")
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

func writeRustDevContainer(buffer *bytes.Buffer) {
	buffer.WriteString(fmt.Sprintf("rust:%s-%s\n\n", RustVersion, DebianAlias))
	buffer.WriteString("CMD [\"/bin/sh\", \"-c\", \"while true; do sleep 30; done;\"]\n")
}

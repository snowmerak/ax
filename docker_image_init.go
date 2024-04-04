package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func initDockerImage(ctx *cli.Context) error {
	name := ctx.Args().First()

	data := ""
	switch {
	case ctx.Bool("go"):
		data, _ = generateGoImageFile()
	case ctx.Bool("node"):
		data, _ = generateNodeImageFile()
	case ctx.Bool("python"):
		data, _ = generatePythonImageFile()
	case ctx.Bool("jdk"):
		data, _ = generateJdkImageFile()
	default:
		return fmt.Errorf("no language specified. Please specify a language: go, node, or python")
	}

	if err := os.WriteFile(fmt.Sprintf("%s.Dockerfile", name), []byte(data), 0644); err != nil {
		return fmt.Errorf("error writing Dockerfile: %w", err)
	}

	return nil
}

func generateGoImageFile() (string, error) {
	builder := strings.Builder{}

	builder.WriteString("FROM ")
	builder.WriteString("golang:latest\n")
	builder.WriteString("\n")
	builder.WriteString("WORKDIR /app\n")
	builder.WriteString("\n")
	builder.WriteString("COPY go.mod .\n")
	builder.WriteString("COPY go.sum .\n")
	builder.WriteString("\n")
	builder.WriteString("RUN go mod download\n")
	builder.WriteString("\n")
	builder.WriteString("COPY . .\n")
	builder.WriteString("\n")
	builder.WriteString("RUN go build -o main .\n")
	builder.WriteString("\n\n")
	builder.WriteString("FROM ")
	builder.WriteString("alpine:latest\n")
	builder.WriteString("\n")
	builder.WriteString("RUN apk --no-cache add ca-certificates\n")
	builder.WriteString("WORKDIR /root/\n")
	builder.WriteString("COPY --from=0 /app/main .\n")
	builder.WriteString("\n")
	builder.WriteString("CMD [\"./main\"]\n")

	return builder.String(), nil
}

func generateNodeImageFile() (string, error) {
	builder := strings.Builder{}

	builder.WriteString("FROM ")
	builder.WriteString("node:latest\n")
	builder.WriteString("\n")
	builder.WriteString("WORKDIR /app\n")
	builder.WriteString("\n")
	builder.WriteString("COPY package*.json .\n")
	builder.WriteString("\n")
	builder.WriteString("RUN npm ci\n")
	builder.WriteString("\n")
	builder.WriteString("COPY . .\n")
	builder.WriteString("\n")
	builder.WriteString("CMD [\"node\", \"dist/index.js\"]\n")

	return builder.String(), nil
}

func generatePythonImageFile() (string, error) {
	builder := strings.Builder{}

	builder.WriteString("FROM ")
	builder.WriteString("python:latest\n")
	builder.WriteString("\n")
	builder.WriteString("WORKDIR /app\n")
	builder.WriteString("\n")
	builder.WriteString("COPY requirements.txt .\n")
	builder.WriteString("\n")
	builder.WriteString("RUN pip install -r requirements.txt\n")
	builder.WriteString("\n")
	builder.WriteString("COPY . .\n")
	builder.WriteString("\n")
	builder.WriteString("CMD [\"python\", \"src/app.py\"]\n")

	return builder.String(), nil
}

func generateJdkImageFile() (string, error) {
	builder := strings.Builder{}

	builder.WriteString("FROM ")
	builder.WriteString("mcr.microsoft.com/openjdk/jdk:21-ubuntu\n")
	builder.WriteString("\n")
	builder.WriteString("WORKDIR /app\n")
	builder.WriteString("\n")
	builder.WriteString("COPY . .\n")
	builder.WriteString("\n")
	builder.WriteString("CMD [\"java\", \"-jar\", \"app.jar\"]\n")

	return builder.String(), nil
}

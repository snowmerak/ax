# ax

```bash
NAME:
   AX - AX is a simple CLI tool for managing your project

USAGE:
   ax command [command options] [arguments...]

COMMANDS:
   init, i      Initialize a new git repository
   push, u      Push and switch back to the previous branch
   commit, c    Commit changes
   feature, f   Create a new feature branch
   bugfix, b    Create a new bugfix branch
   hotfix, x    Create a new hotfix branch
   proposal, p  Create a new proposal branch
   develop, d   Switch to the develop branch
   staging, s   Switch to the staging branch
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Installation

```bash
go install github.com/snowmerak/ax@latest
```

## Usage

### Initialize a new git repository

```bash
ax init <remote-git-url>
```

### Switch to the develop branch

```bash
ax develop
```

### Switch to the staging branch

```bash
ax staging
```

### Switch feature branch

```bash
ax feature <branch-name>
```

### Switch bugfix branch

```bash
ax bugfix <branch-name>
```

### Switch hotfix branch

```bash
ax hotfix <branch-name>
```

### Switch proposal branch

```bash
ax proposal <branch-name>
```

### Commit changes

```bash
ax commit <message>
```

### Push and switch back to the previous branch

```bash
ax push
```

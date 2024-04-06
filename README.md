# ax

```bash
NAME:
   AX - AX is a simple CLI tool for managing your project

USAGE:
   AX [global options] command [command options] 

COMMANDS:
   git, g        
   container, c  
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Installation

```bash
go install github.com/snowmerak/ax@latest
```

## Usage

### git

```
ax g
```

```shell
NAME:
   AX git

USAGE:
   AX git command [command options] 

COMMANDS:
   init, i      Initialize a new git repository
   push, u      Push and switch back to the previous branch
   commit, c    Commit changes
   feature, f   Create a new feature branch
   bugfix, b    Create a new bugfix branch
   hotfix, x    Create a new hotfix branch
   proposal, p  Create a new proposal branch
   unstable, d   Switch to the unstable branch
   stable, s   Switch to the stable branch
   help, h      Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

#### Initialize a new git repository

```bash
ax g init <remote-git-url>
```

Initialize a new git repository and create some branches

1. prod: The production branch, operates on the real server
2. stable: The stable branch, stable version for testing
3. unstable: The development branch, the main branch for development

#### Switch to the stable branch

```bash
ax g stable
```

The stable branch is the stable version for testing.  
You can switch to the stable branch only if you are in the unstable branch.  
If you want to release a new version, you should switch to the stable branch first.

#### Switch to the unstable branch

```bash
ax g unstable
```

The unstable branch is the main branch for development.  
You can switch to the unstable branch only if you are in the stable branch.
If you want to create a new feature, bugfix branch, you should switch to the unstable branch first.  

#### Switch feature branch

```bash
ax g feature <branch-name>
```

The feature branch is used to unstable a new feature.  
You can switch to the feature branch only if you are in the unstable branch.  
If you want to create a new feature branch, you should switch to the unstable branch first.

#### Switch bugfix branch

```bash
ax g bugfix <branch-name>
```

The bugfix branch is used to fix a bug.  
You can switch to the bugfix branch only if you are in the unstable branch.  
If you want to create a new bugfix branch, you should switch to the unstable branch first.  
**CAUTION: The bugfix branch should be merged into the unstable branch before the next release.**

#### Switch hotfix branch

```bash
ax g hotfix <branch-name>
```

The hotfix branch is used to fix a bug on the production server.  
**CAUTION: The hotfix branch must be deleted. Do not merge the hotfix branch into the unstable branch.**

#### Switch proposal branch

```bash
ax g proposal <branch-name>
```

The proposal branch is used to propose a new function or a new idea.  
You can switch to the proposal branch only if you are in the unstable branch.  
If you want to create a new proposal branch, you should switch to the unstable branch first.

#### Commit changes

```bash
ax g commit <message>
```

Commit changes with a message.

#### Push and switch back to the previous branch

```bash
ax g push
```

Push changes and switch back to the previous branch.  
The previous branch is the branch you were in before switching to the current branch.
- If you are in the feature branch, the previous branch is the unstable branch.
- If you are in the bugfix branch, the previous branch is the unstable branch.
- If you are in the hotfix branch, the previous branch is the prod branch.
- If you are in the proposal branch, the previous branch is the feature or bugfix branch.
- If you are in the unstable branch, the previous branch is the stable branch.
- If you are in the stable branch, the previous branch is the prod branch.

### container

```
ax c
```

```shell
NAME:
   AX container

USAGE:
   AX container command [command options] 

COMMANDS:
   init, n   Initialize container config
   image, i  
   dev, d    
   run, r    Run container
   help, h   Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

#### Initialize container config

```shell
ax c n
```

Initialize container config file for build and deployment.

#### Image

```
ax c i
```

```shell
NAME:
   AX container image

USAGE:
   AX container image command [command options] 

COMMANDS:
   init, i   Initialize docker image
   build, b  Build docker image
   help, h   Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

##### Initialize docker image

```shell
ax c i i <dockerfile-name>
```

Initialize docker image file for project building.

You can select a base language for the project.
1. Go: -g or --go
2. Node: -n or --node
3. Python: -p or --python
4. Jdk: -j or --jdk

##### Build docker image

```shell
ax c i b <dockerfile-name>
```

Build docker image.

#### Devcontainer

```
ax c d
```

```shell
NAME:
   AX container dev

USAGE:
   AX container dev command [command options] 

COMMANDS:
   init, i  Initialize dev
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

##### Initialize devcontainer

```shell
ax c d i <devcontainer-name>
```

Initialize devcontainer file for project development.

You can select a base language for the project.
1. Go: -g or --go
2. Node: -n or --node
3. Python: -p or --python
4. Jdk: -j or --jdk
5. C++: -c or --cpp
6. Rust: -r or --rust

#### Run container

```shell
ax c r <alias> <args> ...
```

Run container image with alias and arguments.

- example
- 
```shell
ax c r buf mod init
```

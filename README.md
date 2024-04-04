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
   develop, d   Switch to the develop branch
   staging, s   Switch to the staging branch
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
2. staging: The staging branch, stable version for testing
3. develop: The development branch, the main branch for development

#### Switch to the staging branch

```bash
ax g staging
```

The staging branch is the stable version for testing.  
You can switch to the staging branch only if you are in the develop branch.  
If you want to release a new version, you should switch to the staging branch first.

#### Switch to the develop branch

```bash
ax g develop
```

The develop branch is the main branch for development.  
You can switch to the develop branch only if you are in the staging branch.
If you want to create a new feature, bugfix branch, you should switch to the develop branch first.  

#### Switch feature branch

```bash
ax g feature <branch-name>
```

The feature branch is used to develop a new feature.  
You can switch to the feature branch only if you are in the develop branch.  
If you want to create a new feature branch, you should switch to the develop branch first.

#### Switch bugfix branch

```bash
ax g bugfix <branch-name>
```

The bugfix branch is used to fix a bug.  
You can switch to the bugfix branch only if you are in the develop branch.  
If you want to create a new bugfix branch, you should switch to the develop branch first.  
**CAUTION: The bugfix branch should be merged into the develop branch before the next release.**

#### Switch hotfix branch

```bash
ax g hotfix <branch-name>
```

The hotfix branch is used to fix a bug on the production server.  
**CAUTION: The hotfix branch must be deleted. Do not merge the hotfix branch into the develop branch.**

#### Switch proposal branch

```bash
ax g proposal <branch-name>
```

The proposal branch is used to propose a new function or a new idea.  
You can switch to the proposal branch only if you are in the develop branch.  
If you want to create a new proposal branch, you should switch to the develop branch first.

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
- If you are in the feature branch, the previous branch is the develop branch.
- If you are in the bugfix branch, the previous branch is the develop branch.
- If you are in the hotfix branch, the previous branch is the prod branch.
- If you are in the proposal branch, the previous branch is the feature or bugfix branch.
- If you are in the develop branch, the previous branch is the staging branch.
- If you are in the staging branch, the previous branch is the prod branch.

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
   config, c        Configure container
   image, i         
   devcontainer, d  
   run, r           Run container
   help, h          Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

#### Configure container

```
ax c c
```

```shell
NAME:
   AX container config - Configure container

USAGE:
   AX container config command [command options] 

COMMANDS:
   init, i  Initialize container config
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

##### Initialize container config

```shell
ax c c i
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
   AX container devcontainer

USAGE:
   AX container devcontainer command [command options] 

COMMANDS:
   init, i  Initialize devcontainer
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

```
ax c r
```

```shell
NAME:
   AX container run - Run container

USAGE:
   AX container run command [command options] 

COMMANDS:
   buf, b   Run buf container
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```

##### Run buf container

```shell
ax c r b <args> ...
```

Run buf build cli container with arguments.

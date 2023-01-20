# STACKIT ARGUS CLI

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

CLI for interacting with STACKIT ARGUS API.

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

# Usage

## Installation

Download the desired version for your operating system and processor architecture from the [STACKIT ARGUS CLI](https://github.com/stackitcloud/stackit-argus-cli/releases/tag/v0.0.1). Make the file executable and place it in a directory available in your $PATH.

## Configuration

Set up configurations

```bash
statckit-argus-cli configure
```

After that configuration file(.stackit-argus-cli.yaml) will be generated in your $HOME directory.

## Update Client

The Client will be generated through the OpenAPI Spec file from ARGUS with [openapi-generator.tech](https://openapi-generator.tech/).

To update the generated code, run:

```bash
make generate-client-code
```

## Setup

To get your setup up and running the only thing you have to do is

```bash
make
```

This will initialize a git repo, download the dependencies in the latest versions, install all needed tools and generate binary in ./out/bin directory.
If needed code generation will be triggered in this target as well.

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```

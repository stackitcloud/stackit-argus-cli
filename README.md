# STACKIT ARGUS CLI

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

STACKIT ARGUS CLI & API Client

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

## Update Client

The Client will be generated through the OpenAPI Spec file from ARGUS with [openapi-generator.tech](https://openapi-generator.tech/).

To update the generated code, run:

```bash
make generate-client-code
```

## Setup

To get your setup up and running the only thing you have to do is

```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
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
## Build & run

Build binary

```bash
make build
```

Run cli

```bash
./out/bin/stackit-argus-cli
```
## Configuration

The configuration file location is:

```bash
./cmd/stackit-argus-cli/.stackit-argus-cli.yaml
```

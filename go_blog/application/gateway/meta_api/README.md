<!--
-->

# manager-api

This is a backend project which the dashboard depends on, implemented by Golang.

## Installation

[Please refer to the doc](../README.md)

## Project structure

```text
├── README.md
├── VERSION
├── build-tools
├── build.sh
├── cmd
├── conf
├── entry.sh
├── go.mod
├── go.sum
├── internal
├── run.sh
└── test
```

1. The `cmd` directory is the project entrance.
2. The `internal` directory contains the main logic of manager-api.
3. The `conf` directory contains the default configuration file.
4. The `test` directory contains E2E test cases.

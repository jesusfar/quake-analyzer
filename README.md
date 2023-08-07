# Quake log analyzer

This is a simple implementation of Quake log analyzer, the project follow SOLID principles, 
and is based on [Golang Standar Project Layout] (https://github.com/golang-standards/project-layout).

## Requirements
- Make
- Go >= 20

## How to run

Build:

```shell
make build
```

Run:
```shell
./build/quake-analyzer -report by-death-cause
```

Help:

```shell
./build/quake-analyzer -help
```

## Tests
```shell
make test 
```

## TODO
- Improve log parsing using async process.
- Improve cli interface
- Implement an store for parsed matches.
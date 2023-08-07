# Quake log analyzer

This is a simple implementation of Quake log analyzer, the project follow SOLID principles, 
and is based on [Golang Standar Project Layout](https://github.com/golang-standards/project-layout).

## Architecture

![C4 architecture](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://gist.githubusercontent.com/jesusfar/8ce4dea86c62a9a1fdba012517875eb3/raw/93abca895b019f31c7a9169bd2575e93879d45a6/gistfile1.txt)

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
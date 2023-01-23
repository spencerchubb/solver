# solver

A Rubik's Cube solver written in Go.

This repository has two parts:
- A command line interface (CLI) that can be used on its own
- A Go module that can be used in other Go projects

# Using the CLI

- Download cli/solve.exe
- Open the folder containing solve.exe in a terminal
- Run a command like the following

```bash
./solve -scramble "R U R' F' R U R' U' R' F R2 U' R' U'" -moves UFR -solutions 300
./solve -scramble "R U R' F' R U R' U' R' F R2 U' R' U'" -moves UFDBLR -solutions 10
```

# Using the Go module

Get latest version
```
go get github.com/spencerchubb/solver
```

Get specific version (for example, v0.0.1)
```
go get github.com/spencerchubb/solver@v0.0.1
```

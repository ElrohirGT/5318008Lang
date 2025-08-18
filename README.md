# 5318008Lang

[![Go](https://github.com/ElrohirGT/5318008Lang/actions/workflows/go.yml/badge.svg)](https://github.com/ElrohirGT/5318008Lang/actions/workflows/go.yml)

Don't understand the reference?
[Watch this](https://www.youtube.com/watch?v=r4w2XUqxcBk).

## How to develop this project

Honestly? You only need a Go compiler. But if you need anything else, we
recommend [Nix](https://nixos.org/download/) and
[Flakes](https://nixos.wiki/wiki/flakes)! Then you can simply run:

```bash
nix develop # Enters a devshell with a go compiler and other goodies (like a debugger)
```

## Testing

Almost all project tests are defined under the directory `./tests/`. Every file
you find here is divided in two:

```ts
// Basic CPS file
let a: string = "Hola";
---
// Expected compiler output
```

When you execute `go test` it executes a single test that reads all these files
and compares the output of the compiler with the expected output of the file, if
this differ in any way it reports an error. Easy way to get a big pile of tests
going.

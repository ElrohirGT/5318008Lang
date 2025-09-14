<div>
    <h1 align="center"> 5318008Lang ⚛️</h1>
    <h3 align="center"> 
        Don't understand the reference? 
        <a href="https://www.youtube.com/watch?v=r4w2XUqxcBk">Watch this</a>
    </h3>
</div>

[![Go](https://github.com/ElrohirGT/5318008Lang/actions/workflows/go.yml/badge.svg)](https://github.com/ElrohirGT/5318008Lang/actions/workflows/go.yml)
![GoReport](https://goreportcard.com/badge/github.com/ElrohirGT/5318008Lang)

## How to develop this project

Honestly? You only need a Go compiler. But if you need anything else, we
recommend [Nix](https://nixos.org/download/) and
[Flakes](https://nixos.wiki/wiki/flakes)! Then you can simply run:

```bash
nix develop # Enters a devshell with a go compiler and other goodies (like a debugger)
```

Once you have your environment ready, you can run the project with:

```bash
go run .

go test ./...
```

**Want to change the grammar?**

For that, you should have [antlr4](https://github.com/antlr/antlr4) installed.
Once its ready, modify the file `compiscript/program/Compiscript.g4` and run
this command in the root of the repo:

```bash
antlr4 -Dlanguage=Go -Xexact-output-dir -o ./parser compiscript/program/Compiscript.g4
```

### Fuzzing

Fuzzing has been implemented, checkout for fuzz tests under the
`./tests/fuzzTests/` folder. You can only run one test at a time with:

```bash
go test ./tests/fuzzTests/random_bytes/ -fuzz=Fuzz
```

Specifying the folder test directory each time! Make sure to run them from the
root of the repository.

## 📘 Project Structure

```
.
├── compiscript     // Gramar language definition
├── parser          // Parser module (Generated with ANTLR)
├── listener        // The CORE
├── lib             // Utility function an types
├── tests           // Tests for the diff
└── gui             // Awesome GUI ✨
```

## 🧪 Testing

Almost all project tests are defined under the directory `./tests/`. Every file
you find here is divided in two:

```
// Basic CPS file
let a: string = "Hola";
---
// Expected compiler output
```

When you execute `go test` it executes a single test that reads all these files
and compares the output of the compiler with the expected output defined below
the `---`, if this differ in any way it reports an error. Easy way to get a big
pile of tests going.

### Linting

To lint the project simply run:

```bash
golangci-lint run ./...
```

## ✨ Documentation

### Semantic Analysis

This phase of the code follows the lexical which mean: "The code follows the
desired structure, but does it actually make sense?" In this phase we check
primarily for correctnes in 2 aspects:

- **Types**
- **Scopes**

But the check for those we first need to define how our program represents them,
and interact with the code. In this phase this are the primary modules:

![image](./media/semantic1.png)

#### Parsing

![image](./media/semantic2.png)

#### Types

![image](./media/semantic3.png)

https://github.com/ElrohirGT/5318008Lang/blob/2d8a431ba2d278083dc29979e8506ecb80713d71/listener/types.go#L14-L38

#### Scopes

![image](./media/semantic4.png)

https://github.com/ElrohirGT/5318008Lang/blob/0c25ba6a11598954a7c940915e75ad4d569bfbd9/listener/scopes.go#L45-L62

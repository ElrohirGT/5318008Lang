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

Once you have your environment ready, you can run the project with : 

```bash
go run .

go test
```

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

```ts
// Basic CPS file
let a: string = "Hola";
---
// Expected compiler output
```

When you execute `go test` it executes a single test that reads all these files
and compares the output of the compiler with the expected output defined below the `---`, if
this differ in any way it reports an error. Easy way to get a big pile of tests
going.

## ✨ Documentation 

### Semantic Analysis
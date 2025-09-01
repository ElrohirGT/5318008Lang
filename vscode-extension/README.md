# 531800 Lang - Compiscript Official Extension

Official extension of Visual Studio Code for COmpiscript language. This extension includes support for syntx highlighting, snippets and semantic analysis in real time.

## Caracteristics
- **Syntax highlighting** for `.cps` files
- Predefined **snippets** for common language structures
- **Diagnosis Support**:
  - Syntax errors
  - Semantic errors
- **Integrated command** for running the compiler.

## Ussage
1. Open a file with the `.cps`extension. VSCode will regconize it as Compiscript and will apply the sintax and semantic analysis
2. All the code analysis runs in real time but to see complete compilation logs you can use the command **Run Compiscript File** from the command pallete.
3. In a new terminal will appear the logs from the semantic analysis and the possible errors the code could have.

## Real Time Diagnosis

This is the format for the semantic errors:
```
Error: (line: X, column Y-Z) message
```
And for the syntactic errors:
```
line X:Y message
```

## Included Snippets
Snippets available for teh compiscript language:
- `let`
- `var`
- `const`
- `function`
- `class` 
- `if`
- `while`
- `do while` 
- `for`
- `foreach`
- `print`
- `return`



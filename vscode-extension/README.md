# 531800 Lang - Compiscript Official Extension

Official extension of Visual Studio Code for COmpiscript language. This extension includes support for syntx highlighting, snippets and semantic analysis in real time.

## Caracteristics
- **Syntax highlighting** for `.cps` files
- Predefined **snippets** for common language structures
- **Diagnosis Support**:
  - Syntax errors
  - Semantic errors
- **Integrated command** for running the compiler.

## Usage
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

Snippets available for the Compiscript language:

* **`let`**

  ```cps
  let identifier = value;
  ```
* **`var`**

  ```cps
  var identifier = value;
  ```
* **`const`**

  ```cps
  const identifier = value;
  ```
* **`function`**

  ```cps
  function name(params) {
      // body
  }
  ```
* **`class`**

  ```cps
  class ClassName {
      function constructor(params) {
          // init
      }
  }
  ```
* **`if`**

  ```cps
  if (condition) {
      // body
  }
  ```
* **`while`**

  ```cps
  while (condition) {
      // body
  }
  ```
* **`do while`**

  ```cps
  do {
      // body
  } while (condition);
  ```
* **`for`**

  ```cps
  for (let i = 0; i < n; i++) {
      // body
  }
  ```
* **`foreach`**

  ```cps
  foreach (item in collection) {
      // body
  }
  ```
* **`print`**

  ```cps
  print(expression);
  ```
* **`return`**

  ```cps
  return value;
  ```


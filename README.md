# GO

## Features

- Go is a **_statically typed_** language (have to declare variable types explicitly or they have to be inferred and these types cannot be changed after declaration)

  here we have to specify the type of the variable when creating it

  ```go
  var myVariable string = "Hello, World!"
  ```

  or set the value of the type right away so it can be inferred right away

  ```go
  var myVariable = "Hello, World!"
  ```

- Go is a **_strongly typed_** language i.e the operation you perform on the variable depends on its type.

  in Go

  ```go
  a = 1
  b = "1"
  c = a + b // this will throw an error
  ```

  is not possible because `a` is an integer and `b` is a string.

  In a **_weakly typed_** language like JavaScript, this would be possible because the type of the variable is not enforced.

  Go being statically and strongly typed language is an advantage since the compiler can do more thorough checking of your code and catch errors at compile time rather than runtime. Also there is better code completion and hinting.

- Go is a **_compiled_** language i.e the code is converted into machine code (binary) before execution that can run as a standalone program (unlike interpreted languages like Python or JavaScript where the code is executed line by line so execution time is slower).

- Compilation time itself is very fast in Go.
- Go has built in support for concurrency and parallelism (Goroutines and Channels).
- Simplicity

  This is the general design philosophy of Go. Syntax is pretty concise and we can do a lot with less lines of code. It also has built in **_garbage collection_** which frees up memory when it is no longer needed.

## Setting up Go

- Go to the website [https://go.dev/doc/install](https://go.dev/doc/install) and download the appropriate version for your operating system.
- Follow the instructions on your computer and wait for the installation to complete.
- To check if it's installed open your terminal and type `go version` and if the version pop ups that means it's installed.

## Projects in Go

### Modules and Packages

A package is just folder that contains collection of Go files.

    - `file_1.go`
    - `file_2.go`
    - `file_3.go`

A collection of these packages is called a module. When we're a intialising a project we're basically initialising a new module.

To start a new project/module cd into your project directory and then use `go mod init <name-of-your-module> or <github-repo-link>`

It will create a `go.mod` file containing the name of the module and the Go version we're using. If we start install other external modules, these will be present here along with the version number.

Remember that every Go file is a part of a package. We identify the package to which it belongs to by typing the `package` at the top of the file and then the name of the package. For example:

```go
package myPackage
```

This has to be same for all the files within this folder.

We can also use `package main` where `main` is a special package that tells the compiler to look for the entrypoint function here i.e when creating an executable the compiler needs to know where the program should start from and it will look for a function named main within this main package which serves as the first thing that gets executed in the program.

The `main` package requires the `main` function (this only applies to special package main)

### Functions

To create a function in Go, use the `func` keyword followed by the name of the function. The main function does not accept any parameters and we use `{}` to signify where the code for the function goes.

```go
package main

func main(){}
```

To print something we use the `fmt` package built into Go. To import a package we go just under the package name declaration and then use `import` to import the required package with the package name i.e `import "fmt"`. In Go you have to use the imported packages.

```go
package main
import "fmt"

func main(){
  fmt.Println("Hello World!")
}
```

## Running a Go program

- cd into project directory where the `go.mod` file is present
- compile your program

      `go build <file-location>`
      ```bash
      go build package-1/main.go
      ```
      This produces a binary file called main which we can run to get our output.
      ```bash
      ./main.exe
      ```

- we can directly do it using a single command i.e `go run <file-location>`

```bash
go run package-1/main.go
```

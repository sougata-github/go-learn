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

    file_1.go
    file_2.go
    file_3.go

A collection of these packages is called a module. When we're a intialising a project we're basically initialising a new module.

To start a new project/module cd into your project directory and then use `go mod init <name-of-your-module> or <github-repo-link>`

It will create a `go.mod` file containing the name of the module and the Go version we're using. If we start installing other external modules, these will be present here along with the version number.

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

This produces a binary file called ``main.exe` which we can run to get the output.

```bash
./main.exe
```

- we can directly do it using a single command i.e `go run <file-location>`

```bash
go run package-1/main.go
```

## Variables, Constants and Basic Data Types

### Integers

To declare a variable in Go use the `var` keyword followed by the variable name and the type.

```go
var intNum int
```

Here the `intNum` is typed to `integer` data type. Similar to imports in Go, variables that have been declared have to be used so that there are no unused variables hanging around in the code.

In addition to the `int` type there are

- int8
- int16
- int32
- int64

These represent the memory we use to store our number. 64 can store 4 times larger numbers than 32 but take up more memory. The largest positive number that can be stored in int16 is `32767` — anything larger than this cannot be stored and will throw `overflow` error. However the compiler won't throw any runtime errors, the program will still run but produce weird errors.

```go
func main() {
var intNum int16 = 32767
intNum = intNum+1
fmt.Println(intNum)
}
```

NOTE: `int` will default to 32 bits or 64 bits depending on your system architecture.

In Go we also have access to unsigned ints i.e `unint` with same sizes as ints but only store positive integer so using `unint` we can store +ve integers twice as large in the same amount of memory.

- unint8
- unint16
- unint32
- unint64

int8: `(-128,127)`
unint8: `(0,255)`

For example if we want to store 256 RGB then unsigned 8 bit integer is the best fit rather than 64 bit integer.

### Float

It is used to store decimal numbers. `float64` can store the largest and most precise decimal numbers but they take more memory.

Go doesn't have just a `float` type, you have to specify the number of bits i.e either `float32` or `float64`

We can perform arithmetic operations in Go like `+,-,/ and *` but

- you can't perform operations on mixed types, for example adding an `int32` with `float32` or multiply `int64` with `float64` together. If you need to do it, then you would have to type cast on the values to common type and then proceed

```go
var floatNum32 float32 = 10.1
var intNum32 int32 = 2
var result = floatNum32 + float(intNum32)
```

### Strings

We use a `` `string` `` type to store strings. We can initialise strings with `""` or `` ` ` ``.

```go
var str1 string = "Hello World"
var str2 string = `Hello World`
```

With backticks you format the string directly without having to use escape characters.

```go
var str1 string = "Hello \nWorld"
var str2 string = `Hello
World`
```

We can concatenate strings by adding them

```go
var str1 string = "Hello"
var str2 string = "World"
var result = str1 + " " + str2
```

We can get the length of a string using the built in `len` function but this is not the length of the characters, it's the number of bytes. Since Go uses UTF-8 encoding, characters outside the vanilla ASCII character set are stored with more than a single byte. So for example taking length of ASCII character `A` will result in `1` and taking length of ASCII character `γ` will result in `2`.

For getting the Unicode code points in Go, import the built in `unicode/utf8` package and call the `RuneCountInString` function.

Runes are also a data type in Go that represents characters unicode. A `rune` is basically an alias for int32.

```go
fmt.Println(utf8.RuneCountInString("γ"))
```

To declare a rune, use the rune type and wrap the value in `''`.

```go
var char rune = 'a'
```

### Booleans

Boolean values can be either `true` of `false`

```go
var myBoolean bool = true
```

### Variable Declaration

Up until now we had always initialised our variable whenever we declared it but this is not required. We could create an int variable and then not set it. In these cases Go sets default values depending on its type.

Default values for all integers, unsigned integers, floats, decimals and runes is `0`. For strings it is empty string i.e `""` and for booleans it is `false`.

When we create a variable, we can omit the type and set the variable right away. This way the type is inferred. We could even go a step further and drop the `var` keyword and use the shorthand `:=`.

```go
var myVar = 1
```

```go
myVar := "text"
```

We can initialise multiple variables at once too.

```go
var1, var2 := 1,2
```

For simple types omitting the types when declaring the variales is fine but adding the type when it isn't obvious (say for a function) is a good practice and will make your code easier to follow.

### Constants

Once created you cannot change value of a constant. We have to initialise a constant when declared.

```go
const pi float32 = 3.14
```

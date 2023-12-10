# Go

## Variables

- _ is the blank identifier and mutes the compile-time error returned by unused variables.

1. Using the **var** keyword
    - var x int = 7
    - var s1 string
        s1 = "Learning Go!"

2. Using the short declaration operator (:=)
    - age := 30
    - *We use short declaration when we know the initial value*


## Go Zero Values:
    - numeric types: 0
    - bool type: false
    - string type: "" (empty string)
    - pointer type: nil

## Comments
```go
func main() {
	var a = 4
	var b = 5.2
    
    // Inline comment | To documenting

    /* Block Comment
	a = int(b)
	fmt.Println(a, b) // short-declaration - this is an inline comment
    */
}
```

## Naming Conventions in Go
*(variables, functions, packages)*

- Names start with a letter or an underscore (_)
- Case Sensitive
- Use the first letters of the words
    - ```var mv int ```// mv -> max value
- User fewer letter in smaller scopes and the complete word in larger scopes
    - ```var packetsReceived int ```// NOT OK, to verbose
    - ```var n int```// OK -> no. of packets received
    - ```var taskDone bool```       // ok in larger scopes
- Global variable name starts with uppercase:
    - ```var Max = 100```
- Packages are given lower case single word names:
    - ```maxValue := 1000```
    - ```writeToDB := true```

## Data types
```> 
    //** NUMERIC TYPES **//
 
    // uint8      the set of all unsigned  8-bit integers (0 to 255)
    // uint16      the set of all unsigned 16-bit integers (0 to 65535)
    // uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
    // uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
 
    // int8        the set of all signed  8-bit integers (-128 to 127)
    // int16       the set of all signed 16-bit integers (-32768 to 32767)
    // int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
    // int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
 
    // uint     either 32 or 64 bits
    // int      same size as uint
 
    // float32     the set of all IEEE-754 32-bit floating-point numbers
    // float64     the set of all IEEE-754 64-bit floating-point numbers
    // complex64   the set of all complex numbers with float32 real and imaginary parts
    // complex128  the set of all complex numbers with float64 real and imaginary parts
 
    // byte        alias for uint8
    // rune        alias for int32
```

## Scopes in GO

- The scope or the lifetime of a variable is the interval of time during which it exists as the program executes.
- A name **cannot** be declared again in the same scope
In go there are 3 Scopes:
1. File Scope
2. Package Scope
3. Block(local) Scope

```go
package main

import "fmt" // file scope

const done = false // package scope

func main() {
    x := 10 // local (block) scope
    fmt.Println(x)
}


```

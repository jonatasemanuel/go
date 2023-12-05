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

Title: Go Type Conversions
Date: 29-Apr-2025

Go is a statically-typed language. This means that the type of a variable is known at compile time.  While this offers benefits like early error detection and improved performance, it sometimes requires explicit conversion between types.  This blog post explains how type conversions work in Go.

**The Basics**

Go uses the `T(v)` syntax for type conversions, where `T` is the target type and `v` is the value to convert. The value `v` must be convertible to type `T`.  Let's see a simple example:

```go
package main

import "fmt"

func main() {
    f := 3.14 // float64
    i := int(f) // Convert float64 to int
    fmt.Println(i) // Output: 3
}

```

In this example, we convert a `float64` variable `f` to an `int` variable `i`. Note that Go truncates the floating-point number; it doesn't round it.

**String Conversions**

Converting to and from strings is a common task.  The `strconv` package provides functions for these conversions:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "1234"
    i, err := strconv.Atoi(s) // Convert string to int
    if err != nil {
        panic(err) // Handle the error appropriately in real applications
    }
    fmt.Println(i) // Output: 1234


    n := 42
    str := strconv.Itoa(n)  // Convert int to string
    fmt.Println(str) // Output: "42"

     f := 3.14159
     strFloat := strconv.FormatFloat(f, 'f', 2, 64) // Convert float64 to string with precision
     fmt.Println(strFloat) // Output: "3.14"



}
```

`strconv.Atoi` converts a string to an integer, and `strconv.Itoa` does the reverse.  `strconv.ParseFloat` converts a string to a `float64`, while `strconv.FormatFloat` formats a `float64` as a string.  Always handle potential errors when using these functions.  `strconv` provides functions for other types too, like `ParseBool` and `FormatBool` for booleans.


**Type Assertions**

Type assertions are used with interfaces.  If you have a variable of an interface type, a type assertion allows you to extract the underlying concrete value.

```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    s, ok := i.(string) // Type assertion
    if ok {
        fmt.Println(s) // Output: hello
    } else {
        fmt.Println("Not a string")
    }
}
```

The `ok` variable is a boolean that indicates whether the type assertion succeeded.  If it's `true`, `s` will hold the string value. Otherwise, `s` will be the zero value of the type (empty string in this case).

**Key Considerations**

* **Loss of Information:** Be mindful of potential data loss during conversions, like when converting from `float64` to `int`.
* **Error Handling:**  Always handle errors, particularly when dealing with string conversions.
* **Type Compatibility:** Ensure the types you're converting between are compatible.


By understanding Go's type conversion mechanisms, you can write more robust and flexible code. Remember to choose the right tool for the job, whether it's a direct type conversion, string conversion function, or a type assertion.

Title: Go Defer
Date: 02-Feb-2025

Go's `defer` keyword provides a unique mechanism to schedule function calls to be executed after the surrounding function completes, regardless of whether the function exits successfully or due to a panic. This post will explore `defer` statements, explaining their behavior and common use cases.

### Understanding `defer`

The `defer` keyword is used to postpone the execution of a function until the surrounding function returns. This is particularly useful for cleanup tasks like closing files, releasing resources, or unlocking mutexes.

```go
package main

import "fmt"
import "os"

func main() {
    f, err := os.Create("defer.txt")
    if err != nil {
        panic(err)
    }

    defer f.Close() // Close the file when main() exits

    fmt.Fprintln(f, "Hello, defer!")
    fmt.Println("File written successfully.")
}

```

In this example, `f.Close()` is deferred.  Even if an error occurs during `fmt.Fprintln()`, the `f.Close()` function will still be executed before `main()` returns, ensuring the file is closed properly.

### LIFO Order

When multiple `defer` statements are present within a function, they are executed in Last-In, First-Out (LIFO) order.  This is analogous to a stack.

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Third")
    defer fmt.Println("Second")
    defer fmt.Println("First")

    fmt.Println("Main function executing")
}
```

Output:

```
Main function executing
First
Second
Third
```

### Arguments and `defer`

Arguments to deferred functions are evaluated when the `defer` statement is encountered, not when the deferred function is finally executed.

```go
package main

import "fmt"

func main() {
    i := 1

    defer fmt.Println(i) // Value of i (1) is evaluated here

    i++
    defer fmt.Println(i) // Value of i (2) is evaluated here

    fmt.Println("Main function executing")
}

```

Output:

```
Main function executing
2
1
```

Even though `i` is incremented after the first `defer`, the first deferred function still prints `1` because the argument was evaluated at the time of the `defer` statement.

### Practical Uses

`defer` simplifies resource management significantly. Consider scenarios like database connections, network sockets, or mutex locks. Using `defer` ensures these resources are released promptly, even if errors occur. It makes the code cleaner and less prone to resource leaks.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex

    mu.Lock()
    defer mu.Unlock() // Ensure mutex is unlocked when function exits

    // Access shared resource protected by the mutex
    fmt.Println("Accessing shared resource")

}

```

In this example, `defer mu.Unlock()` guarantees the mutex is unlocked after the function completes, preventing deadlocks.


By understanding and effectively utilizing `defer`, Go developers can write more robust, cleaner, and efficient code by simplifying resource management and ensuring proper cleanup, even in the face of unexpected errors.

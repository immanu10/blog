Title: Go Methods
Date: 07-Mar-2025

Go methods are a powerful way to associate functions with specific data types (structs). They provide a mechanism for creating more organized and reusable code by encapsulating behavior within the types they operate on.  This differs from regular functions which are defined independently.

Think of methods as functions that belong to a particular type.  They have access to the data fields of the associated type, allowing you to manipulate and interact with the object's internal state directly.

```go
package main

import "fmt"

// Define a struct called Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Define a method `Area` that belongs to Rectangle
// The receiver is denoted by (r *Rectangle)
func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Define another method `Perimeter`
func (r *Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}


func main() {
    // Create an instance of the Rectangle struct
    rect := Rectangle{Width: 5, Height: 10}

    // Call the Area method on the rect instance
    area := rect.Area()
    fmt.Println("Area:", area) // Output: Area: 50

    // Call the Perimeter method
    perimeter := rect.Perimeter()
    fmt.Println("Perimeter:", perimeter) // Output: Perimeter: 30

     // Using pointer receivers allows modification of the original struct
    rect.Width = 7
    fmt.Println("New Area:", rect.Area())
}

```

**Key Aspects of Go Methods:**

* **Receiver:** The receiver is specified within parentheses before the method name.  It designates the type the method is associated with.  `(r *Rectangle)` in the example signifies a pointer receiver, meaning the method can modify the original struct.  Using `(r Rectangle)` would create a value receiver, working on a copy.

* **Pointer vs. Value Receivers:**  Pointer receivers are used when you need to modify the underlying struct data within the method. Value receivers operate on a copy, leaving the original struct unchanged.  Choosing the appropriate receiver type is crucial for correct behavior and efficiency.

* **Calling Methods:** Methods are invoked using the dot notation, similar to accessing struct fields: `rect.Area()`.

* **Encapsulation:** Methods contribute to cleaner code organization by grouping related functions with the data they operate on.  This enhances readability and maintainability.

* **Code Reusability:** Methods make your code more reusable.  You can define methods once and use them repeatedly with different instances of the associated type.


Go methods are a fundamental aspect of object-oriented programming in Go, enabling you to write more structured, maintainable, and reusable code. They are powerful tools for encapsulating logic and working with custom data types efficiently.

Title: Go Maps
Date: 20-Feb-2025

Go provides a built-in map type that implements a hash table.  Maps offer an efficient way to store and retrieve values based on a key.  This post will delve into how to use maps in Go.

**Declaring and Initializing**

You can declare a map using the `map` keyword followed by the key type, then the value type.

```go
// Declare a map with string keys and integer values
var myMap map[string]int

// Initialize an empty map
myMap = make(map[string]int)

// Declare and initialize in one line
myMap := map[string]int{"apple": 1, "banana": 2}

// Another way to declare and initialize a map is to use the make function
anotherMap := make(map[int]string)


```

**Adding and Retrieving Values**

Adding a value to a map is done using bracket notation:

```go
myMap["orange"] = 3
```

Retrieving a value works similarly:

```go
value := myMap["apple"] // value will be 1
```

**Checking for Existence**

When retrieving a value, Go returns the zero value for the value type if the key doesn't exist.  To check if a key truly exists, use the following idiom:

```go
value, ok := myMap["grape"]
if ok {
    // Key exists
    fmt.Println("Value:", value) 
} else {
    // Key doesn't exist
    fmt.Println("Key not found")
}
```

**Iterating Over a Map**

You can iterate through a map using a `for...range` loop:

```go
for key, value := range myMap {
    fmt.Println("Key:", key, "Value:", value)
}
```

**Deleting a Key**

The `delete` built-in function removes a key-value pair from a map:

```go
delete(myMap, "apple")
```


**Example: Word Count**

Here's a more complete example demonstrating a word count program:

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "the quick brown fox jumps over the lazy dog"
	words := strings.Fields(text) // Split the string into words

	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++ // Increment the count for each word
	}

	for word, count := range wordCounts {
		fmt.Println(word, ":", count)
	}
}

```


This example shows how maps provide a convenient and efficient way to count the occurrences of words in a text.  Maps are a powerful tool in Go for handling collections of data with key-value relationships.

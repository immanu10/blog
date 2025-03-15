Title: Go Composition
Date: 15-Mar-2025

Go, unlike other object-oriented languages like Java or C++, doesn't have classes. Instead, it uses structs and interfaces to achieve similar functionalities, leveraging a powerful concept called composition. Composition in Go focuses on building complex types by combining simpler ones, promoting code reusability and flexibility. This differs from inheritance, where types inherit properties and methods from a parent class.

Let's illustrate with an example. Imagine building a system to represent different types of content, such as blogs and news articles.  We can define individual structs for each:

```go
type Author struct {
    Name  string
    Email string
}

type BlogPost struct {
    Title    string
    Content  string
    Author   Author
    Comments []string
}

type NewsArticle struct {
    Headline string
    Body     string
    Author   Author
    Source   string
}
```

Notice how both `BlogPost` and `NewsArticle` include an `Author` struct. This is composition in action. Rather than creating a base class and inheriting from it, we embed the `Author` struct directly, allowing `BlogPost` and `NewsArticle` to have an author without explicit inheritance.

We can extend this idea further. Consider adding publishing information:

```go
type PublicationInfo struct {
    PublishedDate string
    Publisher     string
}

type BlogPost struct {
    Title         string
    Content       string
    Author        Author
    Comments      []string
    PublicationInfo PublicationInfo
}


type NewsArticle struct {
    Headline       string
    Body          string
    Author        Author
    Source        string
    PublicationInfo PublicationInfo
}
```

Now, both types incorporate `Author` and `PublicationInfo`. This illustrates the flexibility of composition. You combine the functionalities you need without the rigidity and potential issues of deep inheritance hierarchies.

Accessing embedded fields is straightforward:

```go
func main() {
    blog := BlogPost{
        Title: "Go Composition",
        Content: "Explaining composition...",
        Author: Author{
            Name:  "John Doe",
            Email: "john.doe@example.com",
        },
        PublicationInfo: PublicationInfo{
            PublishedDate: "2024-05-05",
            Publisher: "Example Blog Site",
        },
    }

    fmt.Println("Author Name:", blog.Author.Name)             // Accessing embedded Author
    fmt.Println("Published Date:", blog.PublicationInfo.PublishedDate) // Accessing embedded PublicationInfo

}
```

Go's composition approach encourages creating small, focused types that can be combined in various ways. This modularity leads to more maintainable, understandable, and reusable code.

Title: TypeScript Template Literal Types
Date: 28-Apr-2025

Template literal types were introduced in TypeScript 4.1.  They allow you to create string literals based on string patterns, combining the flexibility of template literals with the type safety of TypeScript. This powerful feature enables you to define types for strings that follow specific formats, such as URLs, file paths, or even formatted data.

Let's explore how template literal types work with some practical examples.

Imagine you want to define types for URLs related to different sections of your website.  Instead of creating separate string literal types for each URL, you can use a template literal type:

```typescript
type SiteSection = "blog" | "docs" | "community";

type SiteURL = `https://example.com/${SiteSection}`;

const blogURL: SiteURL = "https://example.com/blog"; // Valid
const docsURL: SiteURL = "https://example.com/docs"; // Valid
const invalidURL: SiteURL = "https://google.com";    // Error: Type '"https://google.com"' is not assignable to type 'SiteURL'.
```

In this example, `SiteSection` defines the allowed sections.  The `SiteURL` type uses a template literal to create a type that must conform to the pattern `https://example.com/${SiteSection}`.  This ensures that any variable assigned to `SiteURL` must be a valid URL corresponding to one of the defined sections.

You can also combine template literal types with other type operators like unions and intersections. Consider this scenario:

```typescript
type HTTPMethod = "GET" | "POST";

type APIURL = `https://api.example.com/${string}/${HTTPMethod}`;

const validAPIURL: APIURL = "https://api.example.com/users/GET";  // Valid
const anotherValidAPIURL: APIURL = "https://api.example.com/products/POST"; // Valid
const invalidAPIURL: APIURL = "https://api.example.com/users/PUT"; // Error: Type '"https://api.example.com/users/PUT"' is not assignable to type 'APIURL'.
```

Here, `APIURL` uses `${string}` to allow any string in the middle part of the URL, but restricts the HTTP method to `GET` or `POST`.


Template literal types can be particularly useful with functions, guaranteeing that arguments and return values adhere to specific formats.

```typescript
function createGreeting(name: string, language: "en" | "es"): `Hello, ${string}!` | `Hola, ${string}!` {
  if (language === "en") {
    return `Hello, ${name}!`;
  } else {
    return `Hola, ${name}!`;
  }
}

let greeting: `Hello, ${string}!` | `Hola, ${string}!` = createGreeting("John", "en"); // Valid: Type is 'Hello, John!'
greeting = createGreeting("Maria", "es"); // Valid: Type is 'Hola, Maria!'
```


With template literal types, TypeScript offers a robust way to enforce string patterns, enhancing type safety and improving code maintainability.  They provide a concise and expressive way to define complex string types based on dynamic patterns.

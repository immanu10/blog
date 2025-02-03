Title: HTTP Methods
Date: 03-Feb-2025

HTTP methods (also known as verbs) define the actions to be performed on a resource.  They are the foundation of how clients interact with servers over the web.  Understanding these methods is crucial for building any web application.

Here are some of the most common HTTP methods:

* **GET:** Retrieves data from the specified resource.  GET requests should only retrieve data and have no other effect on the data.  They are safe and idempotent, meaning they can be called multiple times without changing the result.  For example, fetching a webpage or an image.

* **POST:** Submits data to be processed to the specified resource.  POST requests are often used to create new resources.  They are neither safe nor idempotent.  For example, submitting a form or creating a new user account.  Data sent in a POST request is typically included in the request body.

* **PUT:**  Replaces all current representations of the target resource with the uploaded content.  PUT is idempotent.  If you send the same PUT request multiple times, it will always have the same result.  For example, updating an existing user's entire profile.

* **PATCH:** Applies partial modifications to a resource.  Unlike PUT, PATCH only updates the specified fields, leaving other fields untouched. For example, updating only a user's email address.

* **DELETE:** Deletes the specified resource.

* **HEAD:** Similar to GET, but only retrieves the headers, not the actual resource body.  Useful for checking if a resource exists or getting its metadata.

* **OPTIONS:**  Describes the communication options for the target resource.  Used to determine what methods the server supports for a particular URL.

**Example: JavaScript Fetch API**

Here's how you might use these methods with JavaScript's `fetch` API:

```javascript
// GET request
fetch('/users')
  .then(response => response.json())
  .then(data => console.log(data));

// POST request
fetch('/users', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({ name: 'John Doe', email: 'john.doe@example.com' }),
})
  .then(response => response.json())
  .then(data => console.log(data));


// PUT request
fetch('/users/123', {
  method: 'PUT',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({ name: 'Jane Doe', email: 'jane.doe@example.com' }),
})
  .then(response => response.json())
  .then(data => console.log(data));

// DELETE Request
fetch('/users/123', { method: 'DELETE' })
  .then(response => console.log(response.status)); // Check the response status

```


Understanding these fundamental HTTP methods is key to effectively interacting with web services and building robust web applications. Each method serves a specific purpose, contributing to the clear and structured exchange of information between clients and servers.

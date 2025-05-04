Title: HTTP Methods
Date: 04-May-2025

HTTP defines a set of request methods to indicate the desired action to be performed for a given resource.  These methods sometimes called HTTP verbs, communicate the type of operation the client wants to perform on the server. Understanding these methods is fundamental to working with web technologies.

Here are some of the most commonly used HTTP methods:

* **GET:** Retrieves data from the specified resource.  GET requests should only retrieve data and should have no other effect on the data. They are safe and idempotent, meaning they can be called multiple times without changing the result.  Parameters are passed in the URL.

    ```javascript
    fetch('https://api.example.com/users', { method: 'GET' })
      .then(response => response.json())
      .then(data => console.log(data));
    ```

* **POST:** Submits data to be processed to the specified resource.  POST requests are often used to create new resources.  They are neither safe nor idempotent, meaning they can have side effects and calling them multiple times can create duplicate resources. Data is sent in the request body.

    ```javascript
    fetch('https://api.example.com/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name: 'John Doe', email: 'john.doe@example.com' })
    })
      .then(response => response.json())
      .then(data => console.log(data));
    ```

* **PUT:** Updates the specified resource with the provided data. PUT requests are idempotent.  If you call PUT multiple times with the same data, it will have the same effect as calling it once.  Typically, the entire resource is updated.

    ```javascript
    fetch('https://api.example.com/users/123', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name: 'Jane Doe', email: 'jane.doe@example.com' })
    })
      .then(response => response.json())
      .then(data => console.log(data));

    ```

* **PATCH:**  Applies partial modifications to a resource.  Unlike PUT, PATCH only updates the fields provided in the request body, leaving other fields unchanged.  PATCH is not necessarily idempotent.

    ```javascript
    fetch('https://api.example.com/users/123', {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email: 'jane.doe.updated@example.com' })
    })
      .then(response => response.json())
      .then(data => console.log(data));
    ```


* **DELETE:** Deletes the specified resource.  DELETE requests are idempotent. Calling DELETE multiple times on the same resource has the same effect as calling it once (assuming the first call successfully deleted the resource).

    ```javascript
    fetch('https://api.example.com/users/123', { method: 'DELETE' })
      .then(response => response.json())
      .then(data => console.log(data));

    ```

These are the fundamental HTTP methods. Understanding their function and implications is crucial for developing and interacting with web services. Other less common methods include HEAD, OPTIONS, CONNECT, and TRACE.

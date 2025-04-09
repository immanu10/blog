Title: Depth-First Search
Date: 09-Apr-2025

Depth-first search (DFS) is a graph traversal algorithm that explores as far as possible along each branch before backtracking.  Imagine traversing a maze – you'd follow one path until you hit a dead end, then retrace your steps to the last intersection and try a different path.  DFS works similarly.

Here's a breakdown of the algorithm and its implementation in JavaScript:

**Algorithm:**

1. **Start at a designated root node.**
2. **Mark the current node as visited.**
3. **Iterate over the unvisited neighbors of the current node.**
4. **For each unvisited neighbor, recursively call the DFS function.**
5. **If all neighbors of the current node have been visited, backtrack.**

**JavaScript Implementation:**

```javascript
function dfs(graph, startNode, visited = new Set()) {
  console.log(`Visiting node: ${startNode}`);
  visited.add(startNode);

  for (const neighbor of graph[startNode]) {
    if (!visited.has(neighbor)) {
      dfs(graph, neighbor, visited);
    }
  }
}


// Example graph represented as an adjacency list
const graph = {
  'A': ['B', 'C'],
  'B': ['D', 'E'],
  'C': ['F'],
  'D': [],
  'E': ['F'],
  'F': []
};

dfs(graph, 'A');
```

**Explanation:**

* The `graph` is represented as an adjacency list, where keys are nodes and values are arrays of their neighbors.
* `visited` is a Set to keep track of visited nodes, preventing infinite loops in cyclic graphs.
* The `dfs` function recursively calls itself on unvisited neighbors.

**Output of the example:**

```
Visiting node: A
Visiting node: B
Visiting node: D
Visiting node: E
Visiting node: F
Visiting node: C
Visiting node: F
```

**Use Cases:**

DFS has various applications, including:

* **Finding paths:** Determining if a path exists between two nodes.
* **Cycle detection:** Identifying cycles in a graph.
* **Topological sorting:** Ordering nodes in a directed acyclic graph.
* **Solving puzzles:**  Like maze solving or finding connected components.

**Time and Space Complexity:**

* **Time Complexity:** O(V + E), where V is the number of vertices and E is the number of edges. This is because each node and edge is visited at most once.
* **Space Complexity:**  O(V) in the worst case due to the recursion stack (in the case of a skewed tree, the maximum depth of the recursion stack will be V).


This explanation and code example provides a foundation for understanding and implementing depth-first search in your projects.  Remember to adapt the implementation based on your specific graph representation and requirements.

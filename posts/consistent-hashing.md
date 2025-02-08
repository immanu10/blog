Title: Consistent Hashing
Date: 08-Feb-2025

Consistent hashing is a distributed caching technique that addresses the limitations of traditional modulo-based hashing. In traditional hashing, when the number of servers changes, a significant portion of the cached data needs to be redistributed, causing a performance hit. Consistent hashing minimizes this redistribution.

Imagine you have a cluster of cache servers.  With traditional hashing, you might assign a key to a server using the modulo operator: `serverIndex = hash(key) % numServers`.  If `numServers` changes, almost all keys will be assigned to a different server.

Consistent hashing solves this by representing the servers and keys on a virtual ring.  Each server is assigned multiple points on the ring using a hash function.  A key is then mapped to the server whose hash is the first encountered clockwise on the ring.

Here's a simplified JavaScript example illustrating the concept:

```javascript
class ConsistentHashing {
  constructor(numReplicas) {
    this.nodes = new Map();
    this.ring = [];
    this.numReplicas = numReplicas || 3;
  }

  addNode(node) {
    for (let i = 0; i < this.numReplicas; i++) {
      const key = `${node}:${i}`;
      const hash = this.hash(key);
      this.nodes.set(hash, node);
      this.ring.push(hash);
    }
    this.ring.sort((a, b) => a - b); // Keep the ring sorted
  }

  removeNode(node) {
    for (let i = 0; i < this.numReplicas; i++) {
      const key = `${node}:${i}`;
      const hash = this.hash(key);
      this.nodes.delete(hash);
      this.ring = this.ring.filter((h) => h !== hash);
    }
  }

  getNode(key) {
    const hash = this.hash(key);
    const index = this.findNearestNodeIndex(hash);
    return this.nodes.get(this.ring[index]);
  }

  findNearestNodeIndex(hash) {
    let left = 0;
    let right = this.ring.length - 1;

    while (left <= right) {
      const mid = Math.floor((left + right) / 2);
      if (this.ring[mid] === hash) {
        return mid;
      } else if (this.ring[mid] < hash) {
        left = mid + 1;
      } else {
        right = mid - 1;
      }
    }

    // Wrap around if no larger hash is found
    return left % this.ring.length;
  }


  // Simple hash function (replace with a better one in production)
  hash(key) {
    let hash = 0;
    for (let i = 0; i < key.length; i++) {
      hash = (hash << 5) - hash + key.charCodeAt(i);
      hash |= 0; // Convert to 32bit integer
    }
    return hash;
  }
}


const ch = new ConsistentHashing();

ch.addNode("server1");
ch.addNode("server2");
ch.addNode("server3");

console.log(ch.getNode("data1")); // server2 (example)
console.log(ch.getNode("data2")); // server3 (example)

ch.addNode("server4"); // Adding a new server

console.log(ch.getNode("data1")); // server2 (likely still the same, demonstrating consistency)
console.log(ch.getNode("data2")); // server3 (likely still the same)

```

When a server is added or removed, only the keys mapped to that server are affected.  This significantly reduces data movement and improves the overall performance and stability of the caching system. Key features of consistent hashing include minimal data movement during server changes, even distribution of keys, and fault tolerance.

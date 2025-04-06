Title: Hash Tables
Date: 06-Apr-2025

Hash tables are a fundamental data structure that provides efficient key-value storage and retrieval.  They achieve this efficiency by using a hash function to map keys to indices in an array (often called a bucket array or hash table). This allows for near constant-time average complexity for insertion, deletion, and lookup operations.

**How Hash Tables Work:**

1. **Hashing:** A hash function takes a key as input and returns a numerical index. A good hash function distributes keys evenly across the available indices to minimize collisions.

2. **Buckets:** The hash table is essentially an array of "buckets." Each bucket can store one or more key-value pairs.

3. **Collisions:** When two different keys hash to the same index (a collision), there are several strategies to handle them:

    * **Separate Chaining:** Each bucket is a linked list or another data structure that can store multiple key-value pairs.  When a collision occurs, the new key-value pair is simply added to the linked list at that index.

    * **Open Addressing (Linear Probing):** If a collision occurs, the hash table searches for the next available empty slot in the array (e.g., by incrementing the index) and inserts the key-value pair there.  Variations like quadratic probing and double hashing use different probing sequences.

**Example in JavaScript (Separate Chaining):**

```javascript
class HashTable {
  constructor(size) {
    this.table = new Array(size);
    this.size = size;
  }

  hash(key) {
    let hash = 0;
    for (let i = 0; i < key.length; i++) {
      hash = (hash + key.charCodeAt(i) * i) % this.size;
    }
    return hash;
  }

  set(key, value) {
    const index = this.hash(key);
    if (!this.table[index]) {
      this.table[index] = [];
    }
    this.table[index].push({ key, value });
  }

  get(key) {
    const index = this.hash(key);
    if (this.table[index]) {
      for (const entry of this.table[index]) {
        if (entry.key === key) {
          return entry.value;
        }
      }
    }
    return undefined; // Key not found
  }
}

const hashTable = new HashTable(5);
hashTable.set("apple", 1);
hashTable.set("banana", 2);
hashTable.set("grape", 3); // Likely collision with a small table size

console.log(hashTable.get("banana")); // Output: 2
console.log(hashTable.get("orange")); // Output: undefined
```

**Time Complexity:**

* **Average Case:**  Insertion, deletion, and lookup are typically O(1).
* **Worst Case:**  If the hash function performs poorly or there are many collisions, operations can degrade to O(n) in the worst-case scenario (especially with separate chaining where the linked list in a bucket becomes very long), where n is the number of elements in the hash table.  Resizing the hash table and rehashing the elements can help mitigate this.

**Use Cases:**

Hash tables are used extensively in:

* Databases (for indexing)
* Caches
* Symbol tables in compilers and interpreters
* Associative arrays
* Deduplication


Hash tables are a powerful tool when you need fast key-value access. Understanding how they work, including collision handling strategies, is crucial for effectively leveraging their performance benefits.

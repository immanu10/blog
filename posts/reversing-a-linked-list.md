Title: Reversing a Linked List
Date: 03-Jan-2025

A linked list is a linear data structure where elements are stored in nodes. Each node points to the next node in the sequence.  Reversing a linked list is a common interview question and a fundamental operation in list manipulation.  It involves changing the direction of the pointers so that the last node becomes the head and the previous head becomes the tail.

There are several ways to reverse a linked list. One of the most common approaches is the iterative method using three pointers: previous, current, and next.

**Iterative Approach:**

```javascript
class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
  }
}

function reverseLinkedList(head) {
  let prev = null;
  let current = head;
  let next = null;

  while (current !== null) {
    next = current.next; // Store the next node
    current.next = prev;   // Reverse the current node's pointer
    prev = current;      // Move prev to the current node
    current = next;      // Move current to the stored next node
  }

  return prev; // prev is now the new head
}


// Example usage:
const node1 = new Node(1);
const node2 = new Node(2);
const node3 = new Node(3);
const node4 = new Node(4);

node1.next = node2;
node2.next = node3;
node3.next = node4;

let head = node1;
console.log("Original Linked List:");
printLinkedList(head);

head = reverseLinkedList(head);

console.log("\nReversed Linked List:");
printLinkedList(head);


function printLinkedList(head) {
  let current = head;
  while (current !== null) {
    process.stdout.write(current.data + " -> ");
    current = current.next;
  }
  console.log("null");
}

```



**Explanation:**

1. **Initialization:**
   - `prev` is initialized to `null` as the initial previous node of the head is null.
   - `current` points to the head of the linked list.
   - `next` is initialized to `null`.

2. **Iteration:** The `while` loop continues as long as `current` is not `null`.
   - `next = current.next`: Stores the next node in the original list. This ensures we don't lose the reference to the next node when we modify `current.next`.
   - `current.next = prev`: Reverses the pointer of the current node to point to the previous node.
   - `prev = current`: Moves `prev` forward, making the current node the new previous node for the next iteration.
   - `current = next`: Moves `current` forward to the previously stored next node.

3. **Return:** After the loop finishes, `prev` will point to the new head (originally the tail), so we return `prev`.


**Key Points:**

* This iterative approach is generally preferred for its efficiency and simplicity.
* It modifies the list in-place, meaning no extra space is used (constant space complexity).
* Understanding the pointer manipulations is crucial to grasping how the reversal works.
* This method can be easily adapted to other languages.


By understanding this iterative approach, you can confidently tackle linked list reversal problems in your interviews and coding tasks.

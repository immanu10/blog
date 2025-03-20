Title: Paxos Explained
Date: 20-Mar-2025

Paxos is a distributed consensus algorithm used to agree on a single value among a group of unreliable processors.  It's foundational to building fault-tolerant distributed systems.  While notoriously difficult to understand, this post breaks down its core components.

**The Problem Paxos Solves**

Imagine multiple servers need to agree on the next database entry.  Network issues, server crashes, and message delays can make reaching a consistent decision tricky.  Paxos ensures that even with these failures:

*   Only a single value is chosen.
*   If no failures occur, processes learn the chosen value.

**Key Roles**

Paxos defines three roles that processes can dynamically assume:

*   **Proposer:** Proposes values to be agreed upon.
*   **Acceptor:**  Votes on proposed values.
*   **Learner:** Learns the chosen value once consensus is reached.

A single process can hold multiple roles.

**Phases of Paxos**

Paxos operates in two phases:

**Phase 1: Prepare**

1.  A Proposer selects a proposal number `n` (higher than any it has used before).
2.  The Proposer sends a `Prepare(n)` message to a majority of Acceptors.

**Phase 2: Promise**

1.  If an Acceptor receives `Prepare(n)` and `n` is higher than any prepare request it has seen:
    *   It promises not to accept any proposal with a number less than `n`.
    *   If it has already accepted a value, it sends that value and its proposal number back to the Proposer.

**Phase 3: Accept**

1.  If the Proposer receives promises from a majority of Acceptors:
    *   It selects a value `v`.  If any Acceptors sent back previously accepted values, the Proposer chooses the value associated with the highest proposal number.
2.  The Proposer sends `Accept(n, v)` to a majority of Acceptors.

**Phase 4: Accepted**

1.  If an Acceptor receives `Accept(n, v)` and it has not promised to ignore proposals numbered `n`:
    *   It accepts the value `v`.
    *   It sends an `Accepted(n, v)` message to all Learners.

**Example (Simplified)**

Let's say we have three Acceptors (A, B, C) and one Proposer (P). P wants to agree on a value (e.g., "Hello").

1.  **Prepare:** P sends `Prepare(1)` to A, B, and C.
2.  **Promise:** A, B, and C respond with promises, as they haven't seen any prior proposals.
3.  **Accept:** P sends `Accept(1, "Hello")` to A, B, and C.
4.  **Accepted:** A, B, and C accept "Hello" and notify Learners.

**Fault Tolerance**

If a Proposer fails, another can take over with a higher proposal number.  If Acceptors fail, as long as a majority remain, the protocol continues.

**JavaScript-like Pseudocode**

```javascript
// Proposer
function propose(value) {
  let proposalNumber = 1;
  // ... Prepare and Promise phase logic ...
  if (majorityPromised) {
    // ... Accept and Accepted phase logic ...
  }
}

// Acceptor
let promisedNumber = 0;
let acceptedValue = null;
let acceptedNumber = 0;

function receivePrepare(n) {
  if (n > promisedNumber) {
    promisedNumber = n;
    // ... send promise ...
  }
}
// ... other functions ...
```

Paxos ensures safety and liveness under certain failure conditions. While complex, understanding its core phases is crucial for building robust distributed systems.

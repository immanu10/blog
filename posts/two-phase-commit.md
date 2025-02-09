Title: Two-Phase Commit
Date: 09-Feb-2025

Two-Phase Commit (2PC) is a distributed consensus protocol used to ensure atomicity (all-or-nothing execution) of transactions spanning multiple database nodes.  Imagine updating balances across multiple accounts in a distributed banking system. 2PC ensures that either all updates succeed, or none do, preventing inconsistencies.

Here's a breakdown of the two phases:

**Phase 1: Prepare Phase**

1. **Coordinator Selection:** One node acts as the coordinator, leading the transaction.
2. **Prepare Request:** The coordinator sends a "prepare" message to all participating nodes (participants).
3. **Participant Vote:** Each participant performs all necessary actions to execute the transaction *locally* but does *not* commit yet.  If a participant can successfully prepare, it logs the transaction and sends a "vote-commit" message to the coordinator. If it cannot prepare (e.g., due to a constraint violation), it sends a "vote-abort" message.

**Phase 2: Commit/Abort Phase**

1. **Coordinator Decision:**
    * If the coordinator receives "vote-commit" from *all* participants, it logs the commit operation and sends a "commit" message to all participants.
    * If the coordinator receives even one "vote-abort" or doesn't hear back from a participant within a timeout period, it decides to abort.  It logs the abort operation and sends an "abort" message to all participants.
2. **Participant Action:**
    * If a participant receives a "commit" message, it applies the previously prepared transaction and then releases any locks held.
    * If a participant receives an "abort" message, it undoes the prepared transaction and releases locks.

**Example (Conceptual JavaScript):**

```javascript
// Coordinator
function coordinator(participants) {
  // Phase 1: Prepare
  const votes = await Promise.all(participants.map(p => p.prepare()));
  if (votes.every(vote => vote === 'vote-commit')) {
    // Phase 2: Commit
    await Promise.all(participants.map(p => p.commit()));
    console.log("Transaction committed");
  } else {
    // Phase 2: Abort
    await Promise.all(participants.map(p => p.abort()));
    console.log("Transaction aborted");
  }
}

// Participant
class Participant {
  async prepare() {
    // Perform local checks and prepare
    try {
      // ... local preparation ...
      return 'vote-commit';
    } catch (error) {
      return 'vote-abort';
    }
  }

  async commit() {
    // Commit the transaction
    // ... apply changes ...
  }

  async abort() {
    // Abort the transaction
    // ... rollback changes ...
  }
}


const participants = [new Participant(), new Participant(), new Participant()];
coordinator(participants);
```

**Limitations:**

* **Blocking:**  2PC can cause blocking.  If the coordinator fails during phase 2, participants may be blocked indefinitely waiting for a commit/abort decision.
* **Performance:**  The two phases and the coordination overhead can impact performance.

Despite its limitations, 2PC provides a strong guarantee of atomicity in distributed transactions and is used in various database systems.  More sophisticated protocols like Three-Phase Commit address some of 2PC's shortcomings but introduce their own complexities.

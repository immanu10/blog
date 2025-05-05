Title: Distributed Consensus
Date: 05-May-2025

A fundamental problem in distributed systems is achieving consensus among multiple nodes.  This means getting all nodes to agree on a single value, even in the presence of failures like network partitions or node crashes.  Distributed consensus is crucial for various applications, such as:

* **Leader election:** Choosing a single node to coordinate tasks.
* **Distributed transactions:** Ensuring data consistency across multiple databases.
* **State machine replication:** Maintaining a consistent state across a cluster of servers.

Several algorithms address distributed consensus. Let's explore Paxos, a widely known, albeit complex, solution.

**Paxos Explained**

Paxos operates in three phases:

1. **Prepare:** A proposer (a node proposing a value) chooses a unique proposal number and sends a "Prepare" message to a majority of acceptors (nodes that can accept proposed values).

2. **Promise:** If an acceptor receives a "Prepare" message with a number higher than any it has seen, it promises not to accept any proposals with lower numbers and sends back information about any previously accepted values.

3. **Accept:** If the proposer receives promises from a majority of acceptors, it chooses a value (either its proposed value or a previously accepted one). It then sends an "Accept" message containing the chosen value and proposal number.

4. **Accepted:**  If an acceptor receives an "Accept" message, it accepts the value unless it has already promised to ignore proposals with that number or lower.

Here’s a simplified example using JavaScript illustrating message passing:

```javascript
// Acceptor
class Acceptor {
  constructor() {
    this.promisedNumber = 0;
    this.acceptedValue = null;
    this.acceptedNumber = 0;
  }

  handlePrepare(proposalNumber) {
    if (proposalNumber > this.promisedNumber) {
      this.promisedNumber = proposalNumber;
      return { promise: true, acceptedValue: this.acceptedValue, acceptedNumber: this.acceptedNumber };
    }
    return { promise: false };
  }

  handleAccept(proposalNumber, value) {
    if (proposalNumber >= this.promisedNumber) {
      this.acceptedValue = value;
      this.acceptedNumber = proposalNumber;
      return { accepted: true };
    }
    return { accepted: false };
  }
}

// Example usage
const acceptor1 = new Acceptor();
const acceptor2 = new Acceptor();
const acceptor3 = new Acceptor();

console.log(acceptor1.handlePrepare(1)); // { promise: true, acceptedValue: null, acceptedNumber: 0 }
console.log(acceptor1.handleAccept(1, "Value 1")); // { accepted: true }
console.log(acceptor1.handlePrepare(2)); // { promise: true, acceptedValue: 'Value 1', acceptedNumber: 1 }


```

This snippet focuses on the core message exchange within Paxos, leaving out the proposer logic and complex scenarios like multiple proposers.  Real-world implementations involve handling failures, timeouts, and ensuring liveness.

Paxos guarantees safety (only a single value is chosen) and eventual liveness (a value will eventually be chosen if a majority of nodes are functioning).  However, it's notoriously difficult to implement and understand fully. Other consensus algorithms like Raft offer simpler implementations while maintaining similar guarantees.

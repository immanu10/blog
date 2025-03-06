Title: Chandy-Lamport Algorithm
Date: 06-Mar-2025

The Chandy-Lamport algorithm is a distributed snapshot algorithm used to record a consistent global state of a distributed system.  It's crucial for debugging, monitoring, and analyzing the system's behavior without stopping or significantly impacting its operation. This algorithm assumes that messages are delivered reliably and in FIFO order between any two processes in the system.

Here's a breakdown of how it works:

1. **Initiation:** A designated process, or any process needing a snapshot, initiates the process. This initiator process records its own local state and sends a special "marker" message along all its outgoing channels.

2. **Marker Propagation:** Upon receiving a marker message for the first time on a channel:

    * The receiving process records its local state.
    * The receiving process records all subsequent messages received on that channel until it receives a marker message on that channel. These recorded messages are considered "in-transit" during the snapshot.
    * The receiving process then sends a marker message along all its outgoing channels.

3. **Termination:** The snapshot process concludes when every process has received a marker message on all its incoming channels.  Each process then sends its recorded local state and the collected "in-transit" messages to the initiator process.

4. **Global State Assembly:** The initiator process assembles the received local states and in-transit messages to construct a consistent global snapshot of the system.

**Example:**

Let's imagine a system with three processes (P1, P2, P3). P1 initiates the snapshot.

```
// Conceptual representation using JavaScript objects for state and messages

// Initial state
P1: { state: "A", sent: [], received: [] }
P2: { state: "B", sent: [], received: [] }
P3: { state: "C", sent: [], received: [] }

// P1 initiates
P1.state = "A1"; // Record local state
P1.sent.push({ to: "P2", marker: true }); // Send marker to P2
P1.sent.push({ to: "P3", marker: true }); // Send marker to P3

// P2 receives marker from P1
P2.received.push({ from: "P1", marker: true });
P2.state = "B1";  // Record local state
P2.sent.push({ to: "P1", msg: "M1" });  // Sends message M1 to P1 (after marker)
P2.sent.push({ to: "P3", marker: true }); // Send marker to P3

// P3 receives marker from P1
P3.received.push({ from: "P1", marker: true });
P3.state = "C1"; // Record local state
P3.sent.push({ to: "P2", msg: "M2"}); // Sends message M2 to P2 (after marker)

// P1 receives M1 from P2 (after its own marker, so in-transit)
P1.received.push({ from: "P2", msg: "M1"});

// P2 receives M2 from P3 (after its own marker to P3, but before receiving P3's marker, so in-transit for P3)

// P3 receives marker from P2
P3.received.push({ from: "P2", marker: true });

// P2 receives marker from P3
P2.received.push({ from: "P3", marker: true });

// P1 receives marker from P3
P1.received.push({ from: "P3", marker: true });

// All processes have received markers on all channels.
// Snapshot consists of local states (A1, B1, C1) and in-transit messages (M1 for P1, M2 for P3).
```

**Key Considerations:**

* **FIFO Channels:** The Chandy-Lamport algorithm relies on FIFO message delivery.  If messages can be reordered, the snapshot might not be consistent.
* **Complexity:** The algorithm's complexity is relatively low, proportional to the number of channels in the system.

The Chandy-Lamport algorithm provides a practical method for capturing consistent snapshots in distributed systems, enabling valuable insights into their behavior without disrupting normal operations.

Title: Lamport Timestamps
Date: 10-Feb-2025

A Lamport timestamp is a simple numerical counter used to assign logical timestamps to events in a distributed system. Unlike physical clocks which can be difficult to synchronize across multiple machines, Lamport timestamps provide a way to order events based on their causal relationships. This ordering is crucial for ensuring data consistency and correctness in distributed applications.

**The Algorithm:**

1. **Local Counter:** Each process maintains a local counter initialized to 0.

2. **Increment on Event:** Every time an event occurs at a process, the local counter is incremented by 1.

3. **Message Passing:** When a process sends a message, it includes its current timestamp in the message.

4. **Receive and Update:** Upon receiving a message, the recipient process updates its local counter by taking the maximum of its current counter and the received timestamp, and then increments its counter by 1.

**Example in Go:**

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Process struct {
    id          int
    timestamp   int
    messageChan chan Message
    mu          sync.Mutex
}

type Message struct {
    senderID   int
    timestamp   int
    messageContent string
}

func (p *Process) run() {
    for {
        select {
        case msg := <-p.messageChan:
            p.mu.Lock()
            p.timestamp = max(p.timestamp, msg.timestamp) + 1
            fmt.Printf("Process %d received message from %d: %s (Timestamp: %d)\n", p.id, msg.senderID, msg.messageContent, p.timestamp)
            p.mu.Unlock()
        default:
            // Simulate some event happening
            time.Sleep(time.Millisecond * 500)
            p.mu.Lock()
            p.timestamp++
            fmt.Printf("Process %d: Internal Event (Timestamp: %d)\n", p.id, p.timestamp)
            p.mu.Unlock()

            // Simulate sending message to another process
             if p.id == 1{
                p2.messageChan <- Message{senderID: p.id, timestamp: p.timestamp, messageContent: "Hello from P1"}
            }


        }
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

var p1, p2 Process

func main() {
    p1 = Process{id: 1, timestamp: 0, messageChan: make(chan Message)}
    p2 = Process{id: 2, timestamp: 0, messageChan: make(chan Message)}

    go p1.run()
    go p2.run()

    select {} // Keep the program running
}


```

**Limitations:**

Lamport timestamps only capture the *happens-before* relationship between events.  Concurrent events can have the same timestamp.  They do not provide true global time ordering.  For scenarios needing strong consistency guarantees about the actual time of events, more complex mechanisms like vector clocks are necessary.

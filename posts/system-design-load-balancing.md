Title: System Design: Load Balancing
Date: 20-Apr-2025

Load balancing is a critical component in designing scalable and reliable systems. It efficiently distributes incoming network traffic across multiple backend servers (also known as a server pool or server farm). This prevents any single server from becoming overloaded, ensuring high availability, fault tolerance, and responsiveness for users.

**Why Use Load Balancing?**

* **High Availability:** If one server fails, the load balancer redirects traffic to the remaining healthy servers, preventing service interruptions.
* **Scalability:**  Easily scale horizontally by adding more servers to the pool as demand increases. The load balancer automatically incorporates new servers into the traffic distribution.
* **Fault Tolerance:**  Redundancy provided by multiple servers ensures that the system can continue operating even if some servers fail.
* **Improved Performance:** Distributing the load prevents any single server from being overwhelmed, resulting in faster response times and a better user experience.
* **Simplified Management:**  Abstracts the complexity of managing multiple servers behind a single entry point.


**How Load Balancing Works:**

A load balancer sits between the client and the server pool, intercepting incoming requests and distributing them across the available servers based on a chosen algorithm.  The client only interacts with the load balancer's IP address and is unaware of the individual servers behind it.

**Load Balancing Algorithms:**

Several algorithms determine how traffic is distributed. Here are a few common ones:

* **Round Robin:** Distributes requests sequentially across servers. Simple and suitable for evenly distributed loads.
* **Least Connections:** Directs requests to the server with the fewest active connections. Effective when server processing times vary.
* **Weighted Round Robin:** Assigns weights to servers, allowing for uneven distribution based on server capacity or other factors. Useful for heterogeneous server environments.
* **IP Hash:**  Uses the client's IP address to determine the server.  Ensures consistent server assignment for a given client.

**Example Scenario (Conceptual JavaScript):**

```javascript
// Conceptual representation - not actual load balancer code

const servers = [
  { address: 'server1.example.com', weight: 2 },
  { address: 'server2.example.com', weight: 1 },
  { address: 'server3.example.com', weight: 1 },
];

function weightedRoundRobin(servers) {
  let totalWeight = servers.reduce((sum, server) => sum + server.weight, 0);
  let currentWeight = 0;

  return function getNextServer() {
    let serverIndex = -1;
    for (let i = 0; i < servers.length; i++) {
        currentWeight += servers[i].weight
        if(currentWeight >= totalWeight){
          serverIndex = i;
          currentWeight = currentWeight % totalWeight
          break;
        }

    }

    // Handle cases where no servers are available
    return serverIndex === -1 ? null : servers[serverIndex].address;
  };

}


const getNext = weightedRoundRobin(servers);

console.log(getNext()); // server1.example.com
console.log(getNext()); // server1.example.com
console.log(getNext()); // server2.example.com
console.log(getNext()); // server3.example.com
console.log(getNext()); // server1.example.com
console.log(getNext()); // server1.example.com

```


**Types of Load Balancers:**

* **Hardware Load Balancers:** Dedicated physical devices offering high performance and advanced features.
* **Software Load Balancers:** Software solutions running on commodity hardware, providing greater flexibility and cost-effectiveness.  Examples include HAProxy, Nginx, and cloud-based load balancing services.



Load balancing is a crucial aspect of building robust and scalable systems.  Understanding the different algorithms and types of load balancers allows you to choose the best solution for your specific needs.

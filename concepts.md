# THE Why: Theory of HTTP and TCP
- Clients make a request to the server and the server responds to the requests with the requested data
- Servers run software that listens to the incoming requests from the clients
- Go routines are used to serve multiple requests at the same time
- The concurrency enables us to use multiple cores at the same time
- Go routines are lighter weight than operating system threads
- For comparison, in Javascript servers are usually single-threaded
- To handle many requests at the same time, javascript servers usually use an async event loop
- The async event loop basically means that whenever a request has to wait on I/O, the server puts it on pause and does something for a bit
- An example is between a client, node.js server and a database:
        1. The first request is made by the first client
        2. The node server processes the first request
        3. The second request is made by the second client
        4. The server requests for information from the database
        5. Switch to processing the second request while waiting for the databases to respond
        6. Sending a response to the second request
        7. Receiving a response from the database
        8. Sending a response to the first request
- Go servers are great for performance whether the workload is I/O or CPU bound
- Node.js and Express work great fo I/O bound tasks, but struggle with CPU bound

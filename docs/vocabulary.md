# Industry Terms
## 01xServers
- Web Server - a computer that serves data over a network typically over the internet.
- Node.js - a javascript runtime environment that lets developers create servers, web apps, command line tools and scripts.
- Express - a Node.js framework
- JS runtime env - software infrastructure where your code is actually executed
- Single threaded - a program or processor only a single command or task at a time in sequential order
- Async Event Loop - is a mechanism that enables the execution of non blocking operations, allowing a single thread to handle multiple tasks concurrently.

# Building Chirpy
- Handler - a handler responds to a http request by carrying out application logic and writing response headers and bodies
- Servemux(router) - stores mapping between predefined URL paths for your applications and corresponding handlers


## 02xRouting
- Middleware -  type of computer software program that provides services to software applications beyond those available from the operating system.
- **Pattern** - a string that specifies the set of URL paths that should be matched to handle HTTP requests
- **Fixed URL paths** - pattern that exactly matches the URL path.
- **Subtree Paths** - if a pattern ends with a slash it matches all URL paths that have the same prefix
- **Longest Match Wins** - if more than one pattern matches the URL path the longest match is chosen, this allows more specificity thus overriding the more general ones
- **Host Specific Patterns** - allow you to serve different content based on the Host header of the request

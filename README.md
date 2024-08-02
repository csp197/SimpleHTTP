[![progress-banner](https://backend.codecrafters.io/progress/http-server/ffaad628-5f0f-4440-be1f-8a0dffd6b6c9)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is a simple HTTP server, built in Golang. The codebase began as a starter from the 
["Build Your Own HTTP server" Challenge](https://app.codecrafters.io/courses/http-server/overview), from Codecrafters.

The server is a very simple [HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)/1.1 server that uses a TCP connection and is capable of serving multiple clients. It also incorporates the [HTTP request syntax](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html) in some custom endpoints.


**Note**: If you'd like to try this challenge for yourself, then please, head over to [codecrafters.io](https://codecrafters.io)!


```
TODO
- Range Requests
   Add support for HTTP range requests, like Range and Content-Range headers, how to handle partial content requests, and more?

- E-Tag caching
   Implement E-Tag caching, like the ETag header, the If-None-Match header, and other mechanisms for caching HTTP responses

- Pipelining
   Support pipelining, like HTTP/1.1's pipelining feature, as well as the ability to re-use a TCP connection for multiple requests.
   https://en.wikipedia.org/wiki/HTTP_pipelining
```
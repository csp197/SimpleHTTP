[![progress-banner](https://backend.codecrafters.io/progress/http-server/ffaad628-5f0f-4440-be1f-8a0dffd6b6c9)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is a simple HTTP server, built in Golang. The codebase began as a starter from the
["Build Your Own HTTP server" Challenge](https://app.codecrafters.io/courses/http-server/overview), from Codecrafters.

The server is a very simple [HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)/1.1 server that uses a TCP connection and is capable of serving multiple clients. It also incorporates the [HTTP request syntax](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html) in some custom endpoints.

**Note**: If you'd like to try this challenge for yourself, then please, head over to [codecrafters.io](https://codecrafters.io)!

Features:

- Binds to a port `4221`
- Responds `2xx` to supported endpoints and requests and `4xx` to unsupported endpoints and requests
- `GET /echo/{str}` and respond with the `{str}` as the response body
- Additional support for gzip compression using the `/echo/{str}` endpoint and the `Accept-Encoding` header
- `GET /user-agent` to read the `User-Agent` header and respond with the value as the response body
- Support for concurrent HTTP connections
- `GET /files/{filename}` to read a file in server filesystem and respond with the file contents as the response body
- `POST /files/{filename}` to create a file in the server filesystem and respond with a `201` status code

TODO:

- Comprehensive testing :D

- Range Requests
  Add support for HTTP range requests, like Range and Content-Range headers, how to handle partial content requests, and more?

- E-Tag caching
  Implement E-Tag caching, like the ETag header, the If-None-Match header, and other mechanisms for caching HTTP responses

- Pipelining
  Support pipelining, like HTTP/1.1's pipelining feature, as well as the ability to re-use a TCP connection for multiple requests.
  https://en.wikipedia.org/wiki/HTTP_pipelining

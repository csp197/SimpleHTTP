[![progress-banner](https://backend.codecrafters.io/progress/http-server/ffaad628-5f0f-4440-be1f-8a0dffd6b6c9)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is a starting point for Go solutions to the
["Build Your Own HTTP server" Challenge](https://app.codecrafters.io/courses/http-server/overview).

[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) is the
protocol that powers the web. In this challenge, you'll build a HTTP/1.1 server
that is capable of serving multiple clients.

Along the way you'll learn about TCP servers,
[HTTP request syntax](https://www.w3.org/Protocols/rfc2616/rfc2616-sec5.html),
and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# Passing the first stage

The entry point for your HTTP server implementation is in `app/server.go`. Study
and uncomment the relevant code, and push your changes to pass the first stage:

```sh
git add .
git commit -m "pass 1st stage" # any msg
git push origin master
```

Time to move on to the next stage!

# Stage 2 & beyond

Note: This section is for stages 2 and beyond.

1. Ensure you have `go (1.19)` installed locally
1. Run `./your_program.sh` to run your program, which is implemented in
   `app/server.go`.
1. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.


```
TODO:
Range Requests
   In this challenge extension, you'll add support for HTTP range requests to your server.
   Along the way, you'll learn about the Range and Content-Range headers, how to handle partial content requests and more.

E-Tag caching
   In this challenge extension, you'll implement E-Tag caching in your HTTP server.
   Along the way, you'll learn about the ETag header, the If-None-Match header, and how E-Tags are used for caching HTTP response.

Pipelining
   In this challenge extension, you'll extend your HTTP server to support pipelining.
   Along the way, you'll learn about HTTP/1.1's pipelining feature and how a HTTP client can re-use a TCP connection for multiple requests.
   https://en.wikipedia.org/wiki/HTTP_pipelining
```
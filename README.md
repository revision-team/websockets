## Chaotic Web Sockets

The intention of this project is to show the behavior of a `server` that can fail some times and a `proxy` that tries to be stable between the `server` and the `clients`.

## instructions

- run server (server will run on port `8081`)

```go
go run server.go
```

- run proxy (proxy will run on port `8082`)

```go
go run proxy.go
```

- run client (client will connect to `proxy`)

```go
go run client.go
```

At this point we should see that the topics are increasing and that are emited from the `server`, readed in the `proxy`, emited from the `proxy` and readed in the `client`.

We can also appreciate the `HeartBeat` comming from the `proxy`.

If we shutdown the `server`, the expected behavior is that the `client` still gets the heartbeat from the `proxy`, which tries to connect to the `server` in order to keep transmitting the messages.

If we setup the `server` again, we should be able to see the normal flow again...

Note: We are not considering subscriptions for the server. Some messages can be lost if we connect multiple clients.

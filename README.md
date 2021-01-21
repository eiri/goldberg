# Rube Goldberg queue server

Ridiculously over-engineered in-memory queue

## About

Named after [Rube Goldberg machine](https://en.wikipedia.org/wiki/Rube_Goldberg_machine), `goldberg` is ridiculously over-engineered in-memory queue server with RPC interface, supporting FIFO, LIFO and double-ended queues.

### Motivation

This is a hobby project made to refresh my memory on handling concurrency in Go. Not intended to pretty much anything practical.

## Usage

Start server (default port 7275) `./goldberg --server --port 1337`. Then call it with `./goldberg -p 1337 -cmd 'ohai!'`

```
$ ./goldberg -s
2021/01/10 10:01:33 Starting server at port :7275
2021/01/10 10:01:52 New connection
2021/01/10 10:01:52 ohai!
2021/01/10 10:01:52 EOF
2021/01/10 10:01:52 Closing connection

$ ./goldberg -cmd 'ohai!'
2021/01/10 10:01:52 Connecting to :7275
2021/01/10 10:01:52 => ohai!
```

## License

[Apache 2.0](https://github.com/eiri/rube-goldberg-queue/blob/main/LICENSE)

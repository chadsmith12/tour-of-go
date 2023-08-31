# Goroutines

A goroutine is a lightweight thread managed by the Go runtime. You can start a new goroutine like so:

```go
go f(x, y, z)
```

The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address sapce, so access to shared memory must be synchronized. The `sync` package provides useful primitives, althought you won't need them much in Go as there are other primitives.

## Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`

```go
ch <- v // Sends v to channel ch
v := <- ch // receives from ch, and assign value to v.
```

Like maps and slices, channels must be created before use:

```go
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. Thjis allows goroutines to synchronize without explicit locks or condition variables. An example of this could be to sum the numbers in a slice, distributing the work between two goroutines. One both goroutines have completed their computation, it calculates the final result.

## Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to `make` to initialize a buffered channel.

```go
ch := make(chan int, 100)
```

Sends to a buffered channel block only when the buffer is full. Receives block with the buffer is empty.

## Range and Close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression. 

```go
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed. The loop `for i := range c` receives values from the channel repeatedly until it is closed.
If you send on a closed channel it will cause a panic.

**NOTE**: Only the sender should close a channel, never the receiver. Sending on a closed channel willl cause a panic.

**Another Note**: Channels aren't like files, you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

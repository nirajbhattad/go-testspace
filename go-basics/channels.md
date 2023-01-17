# Channels


In Go, channels are a way for goroutines to communicate with each other and synchronize their execution. 

Channels are implemented as a way to send and receive values between goroutines.


## Buffered Channels
```
In Go, a buffered channel is a channel that has a buffer associated with it. The buffer is used to store values 
that are sent to the channel before they are received by another goroutine. 

This means that a goroutine can send a value to a buffered channel without being blocked, even if 
there is no other goroutine ready to receive the value.
```

Here's an example that illustrates the use of a buffered channel:

```go

package main

func main(){
    c := make(chan int, 2)
    c <- 1
    c <- 2
    fmt.Println(<-c)
    fmt.Println(<-c)  
}
```

In this example, we create a buffered channel c with a buffer size of 2. We then send the values 1 and 2 to the channel using the c <- 1 and c <- 2 syntax.
Since the buffer size is 2, both values are added to the buffer and the goroutines continue their execution without being blocked. 
The main function then receives the values from the channel using the <-c syntax and prints them.

Application: 

- Buffered channels are useful when you want to improve the performance of your program by reducing the number of times that goroutines are blocked. 
- Buffered channels are also useful when you want to decouple the sending and receiving goroutines, so that the sending goroutine does not need to wait for the 
receiving goroutine to be ready.
- As buffered channels have a buffer, the sending goroutine doesn't need to wait for the receiving goroutine to be ready, 
this allows the program to continue execution while the buffer stores the sent values, improving the performance of the program.


## UnBuffered Channels
```
An unbuffered channel, does not have a buffer associated with it. A goroutine will be blocked when sending a value to an unbuffered channel 
until another goroutine is ready to receive the value. 

A goroutine will also be blocked when trying to receive a value from an unbuffered channel until another goroutine sends a value to the channel.
```

Here's an example that illustrates the use of a unbuffered channel:

```go

package main

func main(){
    c := make(chan int)
    go func(){
      c <- 1
    }()
    fmt.Println(<-c) 
}
```

In this example, we create an unbuffered channel c, and a goroutine that sends the value 1 through the channel. 

When the main function receives the value from the channel using the <-c syntax, it blocks until the goroutine 
has sent a value through the channel.

Application: 

- Unbuffered channels are useful when you need to synchronize the execution of goroutines, and you want to make sure that the sending
  goroutine is blocked until the receiving goroutine is ready to receive the value. 
- It's also useful when you need to make sure that the receiving goroutine is blocked until the sending goroutine has sent a value through the channel.
- As unbuffered channel does not have a buffer, it's mainly used for synchronizing the execution of goroutines and avoid race conditions.



# Concurrency In Go
In Go, concurrency is built into the language through the use of goroutines and channels.

## goroutine

- A goroutine is a lightweight thread of execution. 
- It can be thought of as a function that runs concurrently with other functions.
- A goroutine can be started by simply adding the keyword "go" before a function call. 
- Once started, a goroutine runs independently of the main program flow.

Here's a simple example in Go that demonstrates how a goroutine runs independently of the main program flow:

```go

package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}

func printLetters() {
    for i := 'a'; i <= 'j'; i++ {
        fmt.Printf("%c ", i)
        time.Sleep(time.Second)
    }
}

func main() {
    go printNumbers()
    go printLetters()
    time.Sleep(5 * time.Second)
}
```

In this example, the main function starts two goroutines: one that prints the numbers 1 through 10, and another that prints the letters 'a' through 'j'. 

The main function then sleeps for 5 seconds before exiting.

## gochannels

- Channels are a way for goroutines to communicate with each other and synchronize their execution.
- A channel can be thought of as a type of data structure that allows you to send and receive values. 
- We can think of it as a pipe that connects two goroutines. 
- A goroutine can use the "channel <- value" syntax to send a value to a channel. 
- A goroutine can use the "value <- channel" syntax to receive a value from the channel. 

Here's an example:

```go
package main

import "fmt"

func printNumbers(c chan int) {
    for i := 1; i <= 10; i++ {
        c <- i
    }
    close(c)
}

func main() {
    c := make(chan int)
    go printNumbers(c)
    for n := range c {
        fmt.Println(n)
    }
}
```

In this example, the main function starts a goroutine that prints the numbers 1 through 10. The goroutine uses a channel to send the 
numbers back to the main function, which then prints them. The main function uses the range keyword to iterate over the channel, 
which allows it to receive the numbers one at a time.

Since the goroutine runs concurrently with the main function, the numbers are printed as soon as they are sent on the channel,
rather than waiting for the entire loop to complete. This allows the program to take advantage of multiple CPU cores and improve performance.

As the goroutine runs independently and concurrently, it communicates with the main function through a channel, 
which allows it to synchronize its execution and avoid race conditions.

The scheduler in Go is responsible for managing all the goroutines and their execution efficiently and effectively,
by allocating resources, balancing the load and scheduling the execution of the goroutines in the best way possible.

Together, goroutines and channels make it easy for Go programmers to write concurrent code that is clear, 
efficient and easy to reason about.



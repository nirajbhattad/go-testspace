# goroutine

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

# Concurrency 

Concurrency refers to the ability of a program or system to have multiple tasks in progress at the same time,
but not necessarily executing simultaneously. In other words, it's about dealing with multiple tasks by interleaving their execution.

Here is an example of concurrency in Go:

```go

package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Starting first task")
    go firstTask()
    fmt.Println("Starting second task")
    secondTask()
}

func firstTask() {
    time.Sleep(2 * time.Second)
    fmt.Println("First task done")
}

func secondTask() {
    time.Sleep(3 * time.Second)
    fmt.Println("Second task done")
}

```

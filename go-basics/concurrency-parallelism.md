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

The example above starts two tasks: the firstTask and the secondTask. 

The first task is started using the go keyword, which creates a new goroutine and runs the task concurrently with the main function. 

The second task is started synchronously, within the main function.


# Parallelism

Parallelism refers to the ability of a program or system to have multiple tasks executing simultaneously, typically by using multiple processors or cores. 

In simple words, concurrency is about dealing with a lot of things at once, parallelism is about doing a lot of things at once.

Here is an example of parallelism in Go:

```go 

package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.waitgroup
    wg.Add(2)
    go func(){
      fmt.Println("First task done")
      wg.Done()
    }()
    go func(){
      fmt.Println("Second task done")
      wg.Done()
    }()
    wg.Wait()
} 
```

In this example, we are using WaitGroup to wait for the two goroutines to complete. 

The two tasks are started using the go keyword, which creates new goroutines and runs the tasks in parallel.

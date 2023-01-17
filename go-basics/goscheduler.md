# Scheduler

In Go, goroutines are lightweight threads that are managed by the Go runtime. The Go runtime uses a scheduler to manage the execution of goroutines. The scheduler is responsible for allocating CPU time to different goroutines, and it uses a number of algorithms to determine which goroutines should run and when.


- The scheduler uses a **work-stealing algorithm**, which allows goroutines to steal work from other goroutines that have completed their current task. 
- The scheduler also uses a **priority-based algorithm**, which allows goroutines with higher priority to be scheduled before goroutines with lower priority. 
- Additionally, the scheduler uses a **load-balancing algorithm**, which distributes the workload among available CPU cores.

The scheduler also uses a technique called **"Goroutine preemption"** which allows the scheduler to interrupt the execution of a goroutine and schedule another one if it's deemed necessary. This allows the scheduler to prevent a single goroutine from monopolizing the CPU resources.


Here is an example of how the scheduler works for goroutines in Go:

```go

package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("goroutine 1:", i)
            time.Sleep(time.Millisecond * 100)
        }
    }()
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println("goroutine 2:", i)
            time.Sleep(time.Millisecond * 100)
        }
    }()
    time.Sleep(time.Second * 2)
}

```

In this example, two goroutines are created and run concurrently. Each goroutine has a loop that prints out a message and sleeps for 100 milliseconds. The main goroutine also sleeps for 2 seconds, allowing the other goroutines to run.

When the program is executed, the scheduler will allocate **CPU time** to both goroutines, allowing them to run concurrently. The exact order in which the messages are printed will depend on the **scheduling algorithm** used by the scheduler.

In summary, the Go runtime scheduler is responsible for allocating CPU time to different goroutines. The scheduler uses a number of algorithms, such as **work-stealing, priority-based, and load-balancing algorithms,** to determine which goroutines should run and when. The scheduler also uses a technique called "Goroutine preemption" to interrupt the execution of a goroutine and schedule another one if necessary, this allows the scheduler to prevent a single goroutine from monopolizing the CPU resources.

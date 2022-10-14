package gochallenges

import (
	"fmt"
	"sync"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

var m sync.Map

type Lock struct {
	UUID     *uuid.UUID
	LastUsed int64
}

func PlayWithSyncMap() {
	var wg sync.WaitGroup
	// Initialize

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			uniqueId, _ := uuid.NewV4()
			m.Store(j, Lock{
				UUID:     uniqueId,
				LastUsed: time.Now().UTC().UnixNano(),
			})
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Prints Map Contents
	res := make(map[int]Lock)
	m.Range(func(k, v interface{}) bool {
		fmt.Println("Key: ", k.(int))
		fmt.Println("Value: ", v.(Lock))
		res[k.(int)] = v.(Lock)
		return true
	})

	// Prints and Deletes
	for i := 0; i < 5; i++ {
		t, _ := m.LoadAndDelete(i)
		fmt.Println("for loop: ", t)
	}

	fmt.Println("Done.")
}

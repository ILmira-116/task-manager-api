package tasks

import (
	"fmt"
	"sync"
)

var (
	idCounter int64 = 0
	idMutex   sync.Mutex
)

func GenerateSimpleID() string {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
	return fmt.Sprintf("%d", idCounter)
}

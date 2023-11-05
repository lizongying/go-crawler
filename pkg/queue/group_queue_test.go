package queue

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestNewGroupQueue(t *testing.T) {
	queue := NewGroupQueue(2)

	for i := 1; i <= 5; i++ {
		key, _ := uuid.NewUUID()
		for i1 := 1; i1 <= 4; i1++ {
			value := fmt.Sprintf("Value%d%d", i, i1)
			queue.Enqueue(key.String(), value, int64(i1))
		}

		fmt.Println("len", queue.Size(key.String()))
		items := queue.Get(key.String())

		for _, i2 := range items {
			fmt.Println("value", i2.Value(), i2.Priority())
		}
	}

	fmt.Println("len", queue.Size(""))
}

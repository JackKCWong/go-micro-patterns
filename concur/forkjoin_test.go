package concur

import (
	expect "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestForkJoinGroup_ForkJoin(t *testing.T) {
	var g ForkJoinGroup
	exp := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(exp))
	for i, v := range exp {
		i := i
		v := v
		g.Fork(func() {
			time.Sleep(time.Duration(i*100) * time.Millisecond)
			ch <- v
		})
	}

	select {
	case <-time.After(5 * time.Second):
		expect.Fail(t, "timeout waiting for join")
	case <-g.Join(func() { close(ch) }):
		result := make([]int, 0, 5)
		for v := range ch {
			result = append(result, v)
		}

		expect.Equal(t, exp, result)
	}
}

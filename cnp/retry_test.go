package cnp

import (
	"fmt"
	"testing"
	"time"

	expect "github.com/stretchr/testify/require"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func TestRetry(t *testing.T) {
	cnt := &Counter{}
	result, err := Retry(func() (interface{}, error) {
		if cnt.count < 3 {
			cnt.Increment()
			return nil, fmt.Errorf("simulate error")
		}

		return cnt.count, nil
	}, 5)

	expect.Nil(t, err)
	expect.Equal(t, 3, result)
}

func TestRetryWithInterval(t *testing.T) {
	cnt := &Counter{}

	started := time.Now()
	result, err := RetryWithInterval(func() (int, error) {
		if cnt.count < 3 {
			cnt.Increment()
			return 0, fmt.Errorf("simulate error")
		}

		return cnt.count, nil
	}, 5, 20*time.Millisecond)

	ended := time.Now()

	expect.Nil(t, err)
	expect.Equal(t, 3, result)
	expect.Greater(t, ended.Sub(started), time.Duration(60*time.Millisecond))
}

func TestRetryWithBackoff(t *testing.T) {
	cnt := &Counter{}

	started := time.Now()
	result, err := RetryWithBackoff(func() (int, error) {
		if cnt.count < 3 {
			cnt.Increment()
			return 0, fmt.Errorf("simulate error")
		}

		return cnt.count, nil
	}, 5, 20*time.Millisecond)

	ended := time.Now()

	expect.Nil(t, err)
	expect.Equal(t, 3, result)
	expect.Greater(t, ended.Sub(started), time.Duration(140*time.Millisecond))
}

func TestRetryFailed(t *testing.T) {
	cnt := &Counter{}
	result, err := Retry(func() (interface{}, error) {
		cnt.Increment()
		return nil, fmt.Errorf("simulate error")
	}, 5)

	expect.NotNil(t, err)
	expect.Equal(t, 5, cnt.count)
	expect.Zero(t, result)
}

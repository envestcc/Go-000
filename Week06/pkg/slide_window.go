package pkg

import (
	"errors"
	"sync/atomic"
	"time"
)

type slideWindowCounter struct {
	// buckets 数组长度越大，统计数值会越平滑
	buckets []atomic.Value
	// bucketIndex 当前使用的bucket id
	bucketIndex atomic.Value
	// bucketHead 头部索引
	bucketHead atomic.Value
	// windowSize 滑动窗口的大小
	windowSize time.Duration
}

func NewSlideWindowCounter(bucketNum int, windowDuration time.Duration) *slideWindowCounter {

	c := &slideWindowCounter{
		buckets:     make([]atomic.Value, bucketNum),
		bucketIndex: atomic.Value{},
		bucketHead:  atomic.Value{},
		windowSize:  windowDuration,
	}
	c.Reset()
	return c
}

func (c *slideWindowCounter) Reset() {
	for i := range c.buckets {
		c.buckets[i].Store(0)
	}
	c.bucketIndex.Store(0)
	c.bucketHead.Store(0)
}

func (c *slideWindowCounter) run() error {
	if len(c.buckets) <= 0 {
		return errors.New("buckets size can't be less than 0")
	}
	for {
		time.Sleep(c.windowSize / time.Duration(len(c.buckets)))

		nextIndex := c.bucketIndex.Load().(int) + 1
		if nextIndex >= len(c.buckets) {
			nextIndex = 0
		}
		c.buckets[nextIndex].Store(0)
		c.bucketIndex.Store(nextIndex)

		head := c.bucketHead.Load().(int)
		if nextIndex == head {
			head++
			if head >= len(c.buckets) {
				head = 0
			}
			c.bucketHead.Store(head)
		}

	}
}

func (c *slideWindowCounter) Increment() {
	id := c.bucketIndex.Load().(int)
	cnt := c.buckets[id].Load().(int)
	c.buckets[id].Store(cnt + 1)
}

func (c slideWindowCounter) Avg() int {
	head := c.bucketHead.Load().(int)
	tail := c.bucketIndex.Load().(int)
	sum := 0
	cnt := 0
	for i := head; ; i++ {
		if i >= len(c.buckets) {
			i = 0
		}

		cnt++
		sum += c.buckets[i].Load().(int)

		if i == tail {
			break
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}

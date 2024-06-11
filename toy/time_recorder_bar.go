package toy

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type TimeRecorder struct {
	mu         sync.Mutex
	timestamps []time.Time
}

func NewTimeRecorder() *TimeRecorder {
	return &TimeRecorder{
		timestamps: make([]time.Time, 0),
	}
}

func (t *TimeRecorder) Record() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.timestamps = append(t.timestamps, time.Now())
}

func (t *TimeRecorder) Graph(interval time.Duration) {
	t.mu.Lock()
	defer t.mu.Unlock()

	firstTimestamp := t.timestamps[0]
	intervalCounts := make(map[int]int)

	for _, timestamp := range t.timestamps {
		elapsed := timestamp.Sub(firstTimestamp)
		bucket := int(elapsed / interval)
		intervalCounts[bucket]++
	}

	maxCount := 0
	for _, count := range intervalCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	scale := float64(100) / float64(maxCount)
	if scale > 1 {
		scale = 1
	}

	fmt.Println("Time Graph")
	for i := 0; i <= len(intervalCounts); i++ {
		count, ok := intervalCounts[i]
		if !ok {
			count = 0
		}
		scaledCount := int(float64(count) * scale)
		fmt.Printf("%4d-%-4d: %s %d\n", i, i+1, strings.Repeat("|", scaledCount), count)
	}
}

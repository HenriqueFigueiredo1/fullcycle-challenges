package entity

import (
	"fmt"
	"time"
)

type Results struct {
	TotalTime     time.Duration
	TotalRequests uint64
	TotalSuccess  uint64
	TotalFailure  map[string]uint
}

func NewResults(totalTime time.Duration, totalRequests, totalSuccess uint64, totalFailure map[string]uint) *Results {
	return &Results{
		TotalTime:     totalTime,
		TotalRequests: totalRequests,
		TotalSuccess:  totalSuccess,
		TotalFailure:  totalFailure,
	}
}

func (r *Results) PrintReport() {
	fmt.Println("+--------------------------------------------------------+")
	fmt.Println("| Total time:", r.TotalTime)
	fmt.Println("| Total requests:", r.TotalRequests)
	fmt.Println("| Total success:", r.TotalSuccess)
	for key, value := range r.TotalFailure {
		fmt.Println("| Status", key, ":", value)
	}
	fmt.Println("+--------------------------------------------------------+")
}

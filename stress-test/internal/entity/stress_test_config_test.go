package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcurrencyEqualsOne(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 1, 1)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(1), requests)
	assert.Equal(t, uint64(0), remaining)
}

func TestConcurrencyWithEvenValue2Concurrency1(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 1, 2)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(2), requests)
	assert.Equal(t, uint64(0), remaining)
}

func TestConcurrencyWithEvenValue2Concurrency2(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 2, 2)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(1), requests)
	assert.Equal(t, uint64(0), remaining)
}

func TestConcurrencyWithOddValue3Concurrency2(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 2, 3)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(1), requests)
	assert.Equal(t, uint64(1), remaining)
}

func TestConcurrencyWithOddValue7Concurrency2(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 2, 7)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(3), requests)
	assert.Equal(t, uint64(1), remaining)
}

func TestConcurrencyWithOddValue7Concurrency3(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 3, 7)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(2), requests)
	assert.Equal(t, uint64(1), remaining)
}

func TestConcurrencyWithOddValue7001Concurrency3(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 3, 7001)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(2333), requests)
	assert.Equal(t, uint64(2), remaining)
}

func TestConcurrencyWithOddValue15000001Concurrency3(t *testing.T) {
	stc := NewStressTestConfig("https://.../", 99, 15000001)
	requests, remaining := stc.GetRequestsPerConcurrentExecution()

	assert.Equal(t, uint64(151515), requests)
	assert.Equal(t, uint64(16), remaining)
}

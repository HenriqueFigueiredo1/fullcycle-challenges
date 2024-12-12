package usecase

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/HenriqueFigueiredo1/goexpert-desafio-stress-test/internal/entity"
)

type StressTestUseCase struct{}

type execution struct {
	status int
}

func NewExecuteStressTestUseCase() *StressTestUseCase {
	return &StressTestUseCase{}
}

func (uc *StressTestUseCase) Execute(input *entity.StressTestConfig) *entity.Results {
	start := time.Now()
	requestsPerConcurrency, remainderRequests := input.GetRequestsPerConcurrentExecution()

	wg := &sync.WaitGroup{}
	wg.Add(input.TotalRequests)

	resultChan := make(chan *execution, input.Concurrency)
	for i := 0; i < input.Concurrency; i++ {
		go executeRequests(input.Url, requestsPerConcurrency, i, wg, resultChan)
	}
	go executeRequests(input.Url, remainderRequests, -1, wg, resultChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	retVal := entity.NewResults(0, 0, 0, make(map[string]uint))
	for result := range resultChan {
		retVal.TotalRequests = retVal.TotalRequests + 1
		if result.status >= 200 && result.status < 300 {
			retVal.TotalSuccess = retVal.TotalSuccess + 1
		} else {
			key := strconv.Itoa(result.status)
			retVal.TotalFailure[key] = getOrDefault(retVal.TotalFailure, key, 0) + 1
		}
	}

	retVal.TotalTime = time.Since(start)
	return retVal
}

func executeRequests(url string, totalRequests, id int, wg *sync.WaitGroup, stats chan<- *execution) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	for i := 0; i < totalRequests; i++ {
		req, err := client.Get(url)
		if err != nil {
			stats <- &execution{status: http.StatusInternalServerError}
			wg.Done()
			continue
		}

		stats <- &execution{status: req.StatusCode}
		_ = req.Body.Close()
		wg.Done()
	}
}

func getOrDefault(m map[string]uint, key string, defaultValue uint) uint {
	if value, exists := m[key]; exists {
		return value
	}
	return defaultValue
}

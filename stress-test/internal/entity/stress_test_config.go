package entity

type StressTestConfig struct {
	Url           string
	Concurrency   int
	TotalRequests int
}

func NewStressTestConfig(url string, concurrency, requests int) *StressTestConfig {
	return &StressTestConfig{
		Url:           url,
		Concurrency:   concurrency,
		TotalRequests: requests,
	}
}

func (st *StressTestConfig) GetRequestsPerConcurrentExecution() (int, int) {
	if st.Concurrency == 1 {
		return st.TotalRequests, 0
	} else if st.Concurrency > 1 {
		reqs := st.TotalRequests / st.Concurrency
		remaining := st.TotalRequests % st.Concurrency

		return reqs, remaining
	} else {
		return 0, 0
	}
}

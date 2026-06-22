package services

import (
	"crypto/sha256"
	"fmt"
	"sync"
)


type RequestCounter struct {
	mu       sync.Mutex
	counts   map[string]int
	requests map[string]FizzBuzzRequest
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{
		counts:   make(map[string]int),
		requests: make(map[string]FizzBuzzRequest),
	}
}

func (rc *RequestCounter) Increment(req FizzBuzzRequest) int {
	hash := generateRequestHash(&req)
	rc.mu.Lock()
	defer rc.mu.Unlock()
	rc.counts[hash]++
	rc.requests[hash] = req
	return rc.counts[hash]
}

func (rc *RequestCounter) MostUsed() (FizzBuzzRequest, int) {
    rc.mu.Lock()
    defer rc.mu.Unlock()

    var topHash string
    var topCount int
    for hash, count := range rc.counts {
        if count > topCount {
            topCount = count
            topHash = hash
        }
    }

    return rc.requests[topHash], topCount
}

func generateRequestHash(request *FizzBuzzRequest) string {
	data := fmt.Sprintf("%d-%d-%d-%s-%s", request.Num1, request.Num2, request.Limit, request.Str1, request.Str2)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Rate struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Source string  `json:"source"`
}

type AggregatedRate struct {
	Symbol   string  `json:"symbol"`
	Average  float64 `json:"average"`
	Received int     `json:"sources"`
}

func fetchRate(ctx context.Context, symbol, source, url string, ch chan<- Rate, wg *sync.WaitGroup) {
	defer wg.Done()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error in source %s: %v", source, err)
		return
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("error to decode response: %v", err)
		return
	}

	price, ok := result["bid"].(string)
	if !ok {
		return
	}

	var value float64
	fmt.Sscanf(price, "%f", &value)
	ch <- Rate{Symbol: symbol, Price: value, Source: source}
}

func getAggregatedRate(symbol string) AggregatedRate {
	sources := map[string]string{
		"awesomeapi": fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s-BRL", symbol),
		"mfinance":   fmt.Sprintf("https://api.mercadomineiro.com/%s", symbol), // fictício
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ch := make(chan Rate, len(sources))
	wg := sync.WaitGroup{}

	for name, url := range sources {
		wg.Add(1)
		go fetchRate(ctx, symbol, name, url, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum float64
	count := 0
	for rate := range ch {
		sum += rate.Price
		count++
	}

	avg := 0.0
	if count > 0 {
		avg = sum / float64(count)
	}

	return AggregatedRate{Symbol: symbol, Average: avg, Received: count}
}

func handler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "param 'symbol' required", http.StatusBadRequest)
		return
	}

	result := getAggregatedRate(symbol)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/api/v1/rates", handler)
	fmt.Println("Start server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

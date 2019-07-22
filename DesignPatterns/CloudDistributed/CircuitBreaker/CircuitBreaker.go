package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	st.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		log.Printf("Name: %#+v", name)
		log.Printf("From: %#+v", from)
		log.Printf("to: %#+v", to)
	}
	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	var i = 0
	body, err := cb.Execute(func() (interface{}, error) {
		log.Printf("count: %#+v", i)
		i = i + 1
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func main() {

	for {
		Get("http://192.168.2.12:8085")
	}

}

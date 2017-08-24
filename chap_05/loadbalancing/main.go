package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// Strategy is an interface to be implemented  by loadbalancing
type Strategy interface {
	NextEndpoint() url.URL
	SetEndpoints([]url.URL)
}

// RandomStrategy implements Strategy for random endpoint selection
type RandomStrategy struct {
	endpoints []url.URL
}

// NextEndpoint returns an endpoint using a random strategy
func (s *RandomStrategy) NextEndpoint() url.URL {
	source := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(source)

	return s.endpoints[rd.Intn(len(s.endpoints))]
}

// SetEndpoints ...
func (s *RandomStrategy) SetEndpoints(endpoints []url.URL) {
	s.endpoints = endpoints
}

type LoadBalancer struct {
	strategy Strategy
}

func NewLoadBalancer(s Strategy, endpoints []url.URL) *LoadBalancer {
	s.SetEndpoints(endpoints)
	return &LoadBalancer{strategy: s}
}

func (lb *LoadBalancer) GetEndpoint() url.URL {
	return lb.strategy.NextEndpoint()
}

func (lb *LoadBalancer) UpdateEndpoints(urls []url.URL) {
	lb.strategy.SetEndpoints(urls)
}

func main() {
	endpoints := []url.URL{
		url.URL{Host: "www.google.com"},
		url.URL{Host: "www.baidu.com"},
	}
	lb := NewLoadBalancer(&RandomStrategy{}, endpoints)
	fmt.Println(lb.GetEndpoint())
}

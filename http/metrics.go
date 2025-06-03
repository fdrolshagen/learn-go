package http

import (
	"encoding/json"
	"fmt"
	"sync"
)

type MetricRegistry struct {
	mu      sync.Mutex
	metrics map[string]int
}

func CreateMetricRegistry() *MetricRegistry {
	return &MetricRegistry{
		metrics: make(map[string]int),
	}
}

func (m *MetricRegistry) Increment(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.metrics[key]; !ok {
		m.metrics[key] = 0
	}
	m.metrics[key]++
}

func (m *MetricRegistry) getActuatorRoute() Route {
	return Route{
		method: GET,
		path:   "/actuator/metrics",
		handle: m.Handle,
	}
}

func (m *MetricRegistry) Handle(_ Request) (Response, error) {
	body, err := json.Marshal(m.metrics)
	if err != nil {
		return Response{}, fmt.Errorf("could not marshal metrics: %s", err)
	}

	return Response{
		StatusCode:  200,
		Body:        string(body),
		ContentType: "application/json",
	}, nil
}

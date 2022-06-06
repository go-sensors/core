package temperature

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./consumer.go -destination=./mocks/mock_consumer.go -package=mocks

// Consumer is a consumer of temperatures
type Consumer struct {
	sensor   TemperatureSensor
	handlers []Handler
}

// Handler defines a handler of temperatures
type Handler interface {
	// Handles an individual temperature measurement
	HandleTemperature(*units.Temperature) error
}

// NewConsumer returns a new Consumer with the configured handlers
func NewConsumer(sensor TemperatureSensor, handlers ...Handler) *Consumer {
	return &Consumer{
		sensor,
		handlers,
	}
}

// Run begins consuming from the sensor and handling each measurement. Blocks until either an error occurs or the context is completed.
func (c *Consumer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case temperature := <-c.sensor.Temperatures():
			for _, handler := range c.handlers {
				err := handler.HandleTemperature(temperature)
				if err != nil {
					return err
				}
			}
		}
	}
}

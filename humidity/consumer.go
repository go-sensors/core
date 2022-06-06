package humidity

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./consumer.go -destination=./mocks/mock_consumer.go -package=mocks

// Consumer is a consumer of relative humidities
type Consumer struct {
	sensor   RelativeHumiditySensor
	handlers []Handler
}

// Handler defines a handler of relative humidities
type Handler interface {
	// Handles an individual relative humidity measurement
	HandleRelativeHumidity(*units.RelativeHumidity) error
}

// NewConsumer returns a new Consumer with the configured handlers
func NewConsumer(sensor RelativeHumiditySensor, handlers ...Handler) *Consumer {
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
		case relativeHumidity, ok := <-c.sensor.RelativeHumidities():
			if !ok {
				return nil
			}
			for _, handler := range c.handlers {
				err := handler.HandleRelativeHumidity(relativeHumidity)
				if err != nil {
					return err
				}
			}
		}
	}
}

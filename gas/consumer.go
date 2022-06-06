package gas

import "context"

//go:generate mockgen -source=./consumer.go -destination=./mocks/mock_consumer.go -package=mocks

// Consumer is a consumer of gas concentrations
type Consumer struct {
	sensor   GasSensor
	handlers []Handler
}

// Handler defines a handler of gas concentrations
type Handler interface {
	// Handles an individual concentration measurement
	HandleGasConcentration(*Concentration) error
}

// NewConsumer returns a new Consumer with the configured handlers
func NewConsumer(sensor GasSensor, handlers ...Handler) *Consumer {
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
		case concentration := <-c.sensor.Concentrations():
			for _, handler := range c.handlers {
				err := handler.HandleGasConcentration(concentration)
				if err != nil {
					return err
				}
			}
		}
	}
}

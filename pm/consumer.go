package pm

import "context"

//go:generate mockgen -source=./consumer.go -destination=./mocks/mock_consumer.go -package=mocks

// Consumer is a consumer of particulate matter concentrations
type Consumer struct {
	sensor   ParticulateMatterSensor
	handlers []Handler
}

// Handler defines a handler of particulate matter concentrations
type Handler interface {
	// Handles an individual concentration measurement
	HandlePMConcentration(*Concentration) error
}

// NewConsumer returns a new Consumer with the configured handlers
func NewConsumer(sensor ParticulateMatterSensor, handlers ...Handler) *Consumer {
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
				err := handler.HandlePMConcentration(concentration)
				if err != nil {
					return err
				}
			}
		}
	}
}

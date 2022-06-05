package io

import "io"

//go:generate mockgen -source=./port.go -destination=./mocks/mock_port.go -package=mocks

type PortFactory interface {
	Open() (Port, error)
}

type Port interface {
	Close() error
	io.Reader
	io.Writer
}

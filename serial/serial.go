package serial

import (
	"time"
)

type StopBits byte

const (
	Stop1     StopBits = 1
	Stop1Half StopBits = 15
	Stop2     StopBits = 2
)

type Parity byte

const (
	ParityNone  Parity = 'N'
	ParityOdd   Parity = 'O'
	ParityEven  Parity = 'E'
	ParityMark  Parity = 'M' // parity bit is always 1
	ParitySpace Parity = 'S' // parity bit is always 0
)

type SerialPortConfig struct {
	Baud        int
	ReadTimeout time.Duration
	// Size is the number of data bits. If 0, DefaultSize is used.
	Size byte
	// Parity is the bit to use and defaults to ParityNone (no parity bit).
	Parity Parity
	// Number of stop bits to use. Default is 1 (1 stop bit).
	StopBits StopBits
}

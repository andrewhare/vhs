package sink

import "context"

// Sink is a writable location for output.
type Sink interface {
	Init(ctx context.Context)
	Write(interface{}) error
	Flush() error
}

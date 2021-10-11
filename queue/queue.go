package queue

import (
	"context"
	"io"
	"time"

	"github.com/frain-dev/convoy"
)

type Queuer interface {
	io.Closer
	Write(context.Context, *convoy.Message, time.Duration) error
}

type Job struct {
	Err  error           `json:"err"`
	Data *convoy.Message `json:"data"`
}

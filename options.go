package sse

import (
	"io"
	"net/http"
)

type Logger interface {
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Flags() int
	Output(calldepth int, s string) error
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
	Prefix() string
	Print(v ...any)
	Printf(format string, v ...any)
	Println(v ...any)
	SetFlags(flag int)
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	Writer() io.Writer
}

// Options holds server configurations.
type Options struct {
	// RetryInterval change EventSource default retry interval (milliseconds).
	RetryInterval int
	// Headers allow to set custom headers (useful for CORS support).
	Headers map[string]string
	// ChannelNameFunc allow to create custom channel names.
	// Default channel name is the request path.
	ChannelNameFunc func(*http.Request) string
	// All usage logs end up in Logger
	Logger Logger
}

func (opt *Options) hasHeaders() bool {
	return opt.Headers != nil && len(opt.Headers) > 0
}

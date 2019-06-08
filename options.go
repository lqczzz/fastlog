package fastlog

import (
	"time"

	"go.uber.org/zap/zapcore"
)

const (
	defaultLogLevel      = "INFO"
	defaultLogPath       = "log"
	defaultMaxSize       = 100
	defaultMaxBackups    = 10
	defaultMaxAge        = 3
	defaultEnableConsole = false
	defaultEnableCaller  = false
	defaultBufferSize    = 32 * 1024 * 1024
	defaultFlushInterval = 1 * time.Second
)

func newOptions() *Options {
	return &Options{
		level:          defaultLogLevel,
		path:           defaultLogPath,
		maxSize:        defaultMaxSize,
		maxBackups:     defaultMaxBackups,
		maxAge:         defaultMaxAge,
		enableConsole:  defaultEnableConsole,
		enableCaller:   defaultEnableCaller,
		encoderBuilder: zapcore.NewJSONEncoder,
		bufferSize:     defaultBufferSize,
		flushInterval:  defaultFlushInterval,
	}
}

// Options is options for fastlog
type Options struct {
	path           string
	maxSize        int
	maxBackups     int
	maxAge         int
	level          string
	enableConsole  bool
	enableCaller   bool
	encoderBuilder encoderBuilder
	bufferSize     int
	flushInterval  time.Duration
}

// Option is option of options
type Option func(*Options)
type encoderBuilder func(cfg zapcore.EncoderConfig) zapcore.Encoder

// Path is option for setting log path
func Path(p string) Option {
	return func(o *Options) {
		o.path = p
	}
}

// MaxSize is option for setting max size
func MaxSize(m int) Option {
	return func(o *Options) {
		o.maxSize = m
	}
}

// MaxBackups is option for setting max backups
func MaxBackups(m int) Option {
	return func(o *Options) {
		o.maxBackups = m
	}
}

// MaxAge is option for setting max age
func MaxAge(m int) Option {
	return func(o *Options) {
		o.maxAge = m
	}
}

// Level is option for setting log level
func Level(l string) Option {
	return func(o *Options) {
		o.level = l
	}
}

// EnableConsole is option for enabling console output
func EnableConsole(e bool) Option {
	return func(o *Options) {
		o.enableConsole = e
	}
}

// EnableCaller is option for setting enabling caller
func EnableCaller(e bool) Option {
	return func(o *Options) {
		o.enableCaller = e
	}
}

// BufferSize is option for log buffer
func BufferSize(size int) Option {
	return func(o *Options) {
		o.bufferSize = size
	}
}

// FlushInterval is option for flushing log to disk
func FlushInterval(duration time.Duration) Option {
	return func(o *Options) {
		o.flushInterval = duration
	}
}

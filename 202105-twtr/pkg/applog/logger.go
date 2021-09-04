package applog

import (
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

// Logger is app and access loggers
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	With(fields ...zap.Field) Logger
	Sync(fields ...zap.Field) error
}

var (
	// Fbool constructs a Filed with the given key and value.
	Fbool = zap.Bool
	// Fint constructs a Filed with the given key and value.
	Fint = zap.Int
	// Fint64 constructs a Filed with the given key and value.
	Fint64 = zap.Int64
	// Fuint constructs a Filed with the given key and value.
	Fuint = zap.Uint
	// Fuint64 constructs a Filed with the given key and value.
	Fuint64 = zap.Uint64
	// Fuintptr constructs a Filed with the given key and value.
	Fuintptr = zap.Uintptr
	// Ffloat64 constructs a Filed with the given key and value.
	Ffloat64 = zap.Float64
	// Fstring constructs a Filed with the given key and value.
	Fstring = zap.String
	// Fstringer constructs a Filed with the given key and value.
	Fstringer = zap.Stringer
	// Ftime constructs a Filed with the given key and value.
	Ftime = zap.Time
	// Ferror constructs a Filed with the given key and value.
	Ferror = zap.Error
	// Fduration constructs a Filed with the given key and value.
	Fduration = zap.Duration
	// Fobject constructs a Filed with the given key and value.
	Fobject = zap.Object
	// Label constructs a label with the given key and value.
	Label = zapdriver.Label
	// Fany constructs a Filed with the given key and value.
	Fany = zap.Any
)

package logging

import (
	"go.uber.org/zap"
)

type Field = zap.Field

var (
	Object      = zap.Object
	Array       = zap.Array
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	Bools       = zap.Bools
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex128s = zap.Complex128s
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Complex64s  = zap.Complex64s
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float64s    = zap.Float64s
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Float32s    = zap.Float32s
	Int         = zap.Int
	Intp        = zap.Intp
	Ints        = zap.Ints
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int64s      = zap.Int64s
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int32s      = zap.Int32s
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int16s      = zap.Int16s
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	Int8s       = zap.Int8s
	String      = zap.String
	Stringp     = zap.Stringp
	Strings     = zap.Strings
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uints       = zap.Uints
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint64s     = zap.Uint64s
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint32s     = zap.Uint32s
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint16s     = zap.Uint16s
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Binary      = zap.Binary
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Uintptrs    = zap.Uintptrs
	Time        = zap.Time
	Timep       = zap.Timep
	Times       = zap.Times
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Durations   = zap.Durations
	NamedError  = zap.NamedError
	Errors      = zap.Errors
	Stringer    = zap.Stringer
	Reflect     = zap.Reflect
)

package logit

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field zap.Field

func Any(key string, value interface{}) Field {
	return Field(any(key, value))
}

// nolint:gocyclo
func any(key string, value interface{}) zap.Field {
	switch val := value.(type) {
	case zapcore.ObjectMarshaler:
		return zap.Object(key, val)
	case zapcore.ArrayMarshaler:
		return zap.Array(key, val)
	case bool:
		return zap.Bool(key, val)
	case *bool:
		return zap.Boolp(key, val)
	case []bool:
		return zap.Bools(key, val)
	case complex128:
		return zap.Complex128(key, val)
	case *complex128:
		return zap.Complex128p(key, val)
	case []complex128:
		return zap.Complex128s(key, val)
	case complex64:
		return zap.Complex64(key, val)
	case *complex64:
		return zap.Complex64p(key, val)
	case []complex64:
		return zap.Complex64s(key, val)
	case float64:
		return zap.Float64(key, val)
	case *float64:
		return zap.Float64p(key, val)
	case []float64:
		return zap.Float64s(key, val)
	case float32:
		return zap.Float32(key, val)
	case *float32:
		return zap.Float32p(key, val)
	case []float32:
		return zap.Float32s(key, val)
	case int:
		return zap.Int(key, val)
	case *int:
		return zap.Intp(key, val)
	case []int:
		return zap.Ints(key, val)
	case int64:
		return zap.Int64(key, val)
	case *int64:
		return zap.Int64p(key, val)
	case []int64:
		return zap.Int64s(key, val)
	case int32:
		return zap.Int32(key, val)
	case *int32:
		return zap.Int32p(key, val)
	case []int32:
		return zap.Int32s(key, val)
	case int16:
		return zap.Int16(key, val)
	case *int16:
		return zap.Int16p(key, val)
	case []int16:
		return zap.Int16s(key, val)
	case int8:
		return zap.Int8(key, val)
	case *int8:
		return zap.Int8p(key, val)
	case []int8:
		return zap.Int8s(key, val)
	case string:
		return zap.String(key, val)
	case *string:
		return zap.Stringp(key, val)
	case []string:
		return zap.Strings(key, val)
	case uint:
		return zap.Uint(key, val)
	case *uint:
		return zap.Uintp(key, val)
	case []uint:
		return zap.Uints(key, val)
	case uint64:
		return zap.Uint64(key, val)
	case *uint64:
		return zap.Uint64p(key, val)
	case []uint64:
		return zap.Uint64s(key, val)
	case uint32:
		return zap.Uint32(key, val)
	case *uint32:
		return zap.Uint32p(key, val)
	case []uint32:
		return zap.Uint32s(key, val)
	case uint16:
		return zap.Uint16(key, val)
	case *uint16:
		return zap.Uint16p(key, val)
	case []uint16:
		return zap.Uint16s(key, val)
	case uint8:
		return zap.Uint8(key, val)
	case *uint8:
		return zap.Uint8p(key, val)
	case []byte:
		return zap.String(key, string(val))
	case uintptr:
		return zap.Uintptr(key, val)
	case *uintptr:
		return zap.Uintptrp(key, val)
	case []uintptr:
		return zap.Uintptrs(key, val)
	case time.Time:
		return zap.Time(key, val)
	case *time.Time:
		return zap.Timep(key, val)
	case []time.Time:
		return zap.Times(key, val)
	case time.Duration:
		return zap.Duration(key, val)
	case *time.Duration:
		return zap.Durationp(key, val)
	case []time.Duration:
		return zap.Durations(key, val)
	case error:
		return zap.NamedError(key, val)
	case []error:
		return zap.Errors(key, val)
	case fmt.Stringer:
		return zap.Stringer(key, val)
	default:
		j, err := json.Marshal(val)
		if err != nil {
			return zap.Error(err)
		}
		return zap.String(key, string(j))
	}
}

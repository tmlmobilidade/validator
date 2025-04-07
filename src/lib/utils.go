package lib

import (
	"fmt"
	"reflect"
	"strconv"
)

func ParseStringToPrimitive[T any](str string, t *T, errors *[]string) {
	if str == "" {
		return
	}

	switch any(*t).(type) {
	case string:
		*t = any(str).(T)
		return
	case int8:
		f, err := strconv.ParseInt(str, 10, 8)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to int8", str))
		}
		*t = any(int8(f)).(T)
		return
	case int16:
		f, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to int16", str))
		}
		*t = any(int16(f)).(T)
		return
	case int32:
		f, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to int32", str))
		}
		*t = any(int32(f)).(T)
		return
	case int64:
		f, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to int64", str))
		}
		*t = any(int64(f)).(T)
		return
	case int:
		f, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to int", str))
		}
		*t = any(int(f)).(T)
		return
	case uint8:
		f, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to uint8", str))
		}
		*t = any(uint8(f)).(T)
		return
	case uint16:
		f, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to uint16", str))
		}
		*t = any(uint16(f)).(T)
		return
	case uint32:
		f, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to uint32", str))
		}
		*t = any(uint32(f)).(T)
		return
	case uint64:
		f, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to uint64", str))
		}
		*t = any(uint64(f)).(T)
		return
	case uint:
		f, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to uint", str))
		}
		*t = any(uint(f)).(T)
		return
	case float32:
		f, err := strconv.ParseFloat(str, 32)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to float32", str))
		}
		*t = any(float32(f)).(T)
		return
	case float64:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to float64", str))
		}
		*t = any(f).(T)
		return
	case bool:
		f, err := strconv.ParseBool(str)
		if err != nil {
			*errors = append(*errors, fmt.Sprintf("Failed to parse \"%s\" to bool", str))
		}
		*t = any(f).(T)
		return
	default:
		*errors = append(*errors, fmt.Sprintf("unsupported type: %s", reflect.TypeOf(*t)))
	}
}
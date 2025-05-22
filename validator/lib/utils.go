package lib

import (
	"encoding/json"
	"fmt"
	"maps"
	"reflect"
	"strconv"
)

// ParseStringToPrimitive converts a string to a primitive type and stores it in the provided pointer.
// If the string is empty, the function returns without modifying the target value.
// If parsing fails, an error message is appended to the provided errors slice.
//
// The function supports the following primitive types:
//   - string
//   - int8, int16, int32, int64, int
//   - uint8, uint16, uint32, uint64, uint
//   - float32, float64
//   - bool
//
// If an unsupported type is provided, the function will panic with an error message.
//
//	@param str string - The string to parse
//	@param t *T - Pointer to the target variable where the parsed value will be stored
//	@return msg string - The error message if parsing fails, empty string otherwise
func ParseStringToPrimitive[T any](str string, t *T) (msg string) {
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
			return fmt.Sprintf("Failed to parse \"%s\" to int8", str)
		}
		*t = any(int8(f)).(T)
		return
	case int16:
		f, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to int16", str)
		}
		*t = any(int16(f)).(T)
		return
	case int32:
		f, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to int32", str)
		}
		*t = any(int32(f)).(T)
		return
	case int64:
		f, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to int64", str)
		}
		*t = any(int64(f)).(T)
		return
	case int:
		f, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to int", str)
		}
		*t = any(int(f)).(T)
		return
	case uint8:
		f, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to uint8", str)
		}
		*t = any(uint8(f)).(T)
		return
	case uint16:
		f, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to uint16", str)
		}
		*t = any(uint16(f)).(T)
		return
	case uint32:
		f, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to uint32", str)
		}
		*t = any(uint32(f)).(T)
		return
	case uint64:
		f, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to uint64", str)
		}
		*t = any(uint64(f)).(T)
		return
	case uint:
		f, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to uint", str)
		}
		*t = any(uint(f)).(T)
		return
	case float32:
		f, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to float32", str)
		}
		*t = any(float32(f)).(T)
		return
	case float64:
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to float64", str)
		}
		*t = any(f).(T)
		return
	case bool:
		f, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Sprintf("Failed to parse \"%s\" to bool", str)
		}
		*t = any(f).(T)
		return
	default:
		//Panic with error
		panic(fmt.Sprintf("variable \"%s\" is of type \"%s\" and cannot be parsed to primitive type", str, reflect.TypeOf(*t)))
	}
}

// PrintMap takes any value and prints it as indented JSON to stdout.
// If there is an error marshaling the value to JSON, it prints the error.
//
//	@param a any - The value to print as JSON
//	@param minify ...bool - Optional parameter to minify the output
func PrintMap(a any, minify ...bool) {
	shouldMinify := false

	if len(minify) > 0 {
		shouldMinify = minify[0]
	}
	
	if shouldMinify {
		b, err := json.Marshal(a)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%s\n", string(b))
	} else {
		b, err := json.MarshalIndent(a, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%s\n", string(b))
	}
}

// Returns a if condition is true, otherwise returns b
// Substitute for the ternary operator
//
//	@param condition bool - The condition to check
//	@param a T - The value to return if the condition is true
//	@param b T - The value to return if the condition is false
//	@return T - The value to return
func IfThenElse[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// MergeMaps merges two maps into a single map
//
//	@param a map[string]string - The first map
//	@param b map[string]string - The second map
//	@return map[string]string - The merged map
func MergeMaps[T any](a, b map[string]T) map[string]T {
	result := make(map[string]T)
	maps.Copy(result, a)
	maps.Copy(result, b)

	return result
}
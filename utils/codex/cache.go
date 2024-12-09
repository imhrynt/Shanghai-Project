package codex

import (
	"bytes"
	"reflect"
	"sync"
)

var functionCache sync.Map

type (
	reader func(rv reflect.Value)
	writer func(buf *bytes.Buffer, rv reflect.Value)
)

func getWriter(rt reflect.Type) (writer, error) {
	fw := makeWriter(rt)
	return fw, nil
}

func makeWriter(rt reflect.Type) writer {
	rk := rt.Kind()
	switch {
	case rk == reflect.Bool:
		return nil
	case rk >= reflect.Uint8 && rk <= reflect.Uint64:
		return nil
	case rk >= reflect.Int8 && rk <= reflect.Int64:
		return nil
	case rk == reflect.String:
		return nil
	case rk == reflect.Slice:
		return nil
	case rk == reflect.Array:
		return nil
	case rk == reflect.Map:
		return nil
	case rk == reflect.Struct:
		return nil
	case rk == reflect.Interface:
		return nil
	default:
		return nil
	}
}

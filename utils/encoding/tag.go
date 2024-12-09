package encoding

import "reflect"

func isPrimitive(kind reflect.Kind) bool {
	return kind == reflect.Bool || isInt(kind) || isUint(kind) ||
		isFloat(kind) || isComplex(kind) || kind == reflect.String
}

func isInt(kind reflect.Kind) bool {
	return kind >= reflect.Int && kind <= reflect.Int64
}

func isUint(kind reflect.Kind) bool {
	return kind >= reflect.Uint && kind <= reflect.Uintptr
}

func isFloat(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

func isComplex(kind reflect.Kind) bool {
	return kind == reflect.Complex64 || kind == reflect.Complex128
}

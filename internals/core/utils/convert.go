package utils

import (
	"strconv"

	"github.com/HEEPOKE/generate-db/pkg/enums"
)

func ConvertToFloat32(s string) float32 {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	f32 := float32(f64)

	return f32
}

func ConvertToFloat64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f64
}

func ConvertToUint(s string) uint {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(u64)
}

func ConvertToUint8(s string) uint8 {
	u64, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0
	}
	return uint8(u64)
}

func ConvertToUint16(s string) uint16 {
	u64, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(u64)
}

func ConvertToUint32(s string) uint32 {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(u64)
}

func ConvertToColumnType(value string, columnType enums.Category) interface{} {
	switch columnType {
	case enums.FLOAT32:
		return ConvertToFloat32(value)
	case enums.FLOAT64:
		return ConvertToFloat64(value)
	case enums.UINT:
		return ConvertToUint(value)
	case enums.UINT8:
		return ConvertToUint8(value)
	case enums.UINT16:
		return ConvertToUint16(value)
	case enums.UINT32:
		return ConvertToUint32(value)
	default:
		return nil
	}
}

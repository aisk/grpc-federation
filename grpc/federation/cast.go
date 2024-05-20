package federation

import "fmt"

func Int64ToInt8(v int64) (int8, error) {
	if v < -1*(1<<7) || (1<<7) <= v {
		return 0, fmt.Errorf("failed to convert int64(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

func Int64ToInt16(v int64) (int16, error) {
	if v < -1*(1<<15) || (1<<15) <= v {
		return 0, fmt.Errorf("failed to convert int64(%d) to int16: %w", v, ErrOverflowTypeConversion)
	}
	return int16(v), nil
}

func Int64ToInt32(v int64) (int32, error) {
	if v < -1*(1<<31) || (1<<31) <= v {
		return 0, fmt.Errorf("failed to convert int64(%d) to int32: %w", v, ErrOverflowTypeConversion)
	}
	return int32(v), nil
}

func Int64ToUint8(v int64) (uint8, error) {
	if v < 0 || (1<<8) <= v {
		return 0, fmt.Errorf("failed to convert int64(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Int64ToUint16(v int64) (uint16, error) {
	if v < 0 || (1<<16) <= v {
		return 0, fmt.Errorf("failed to convert int64(%d) to uint16: %w", v, ErrOverflowTypeConversion)
	}
	return uint16(v), nil
}

func Int64ToUint32(v int64) (uint32, error) {
	if v < 0 || (1<<32) <= v {
		return 0, fmt.Errorf("failed to convert int64(%d) to uint32: %w", v, ErrOverflowTypeConversion)
	}
	return uint32(v), nil
}

func Int64ToUint64(v int64) (uint64, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int64(%d) to uint64: %w", v, ErrOverflowTypeConversion)
	}
	return uint64(v), nil
}

func Int32ToInt8(v int32) (int8, error) {
	if v < -1*(1<<7) || (1<<7) <= v {
		return 0, fmt.Errorf("failed to convert int32(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

func Int32ToInt16(v int32) (int16, error) {
	if v < -1*(1<<15) || (1<<15) <= v {
		return 0, fmt.Errorf("failed to convert int32(%d) to int16: %w", v, ErrOverflowTypeConversion)
	}
	return int16(v), nil
}

func Int32ToUint8(v int32) (uint8, error) {
	if v < 0 || (1<<8) <= v {
		return 0, fmt.Errorf("failed to convert int32(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Int32ToUint16(v int32) (uint16, error) {
	if v < 0 || (1<<16) <= v {
		return 0, fmt.Errorf("failed to convert int32(%d) to uint16: %w", v, ErrOverflowTypeConversion)
	}
	return uint16(v), nil
}

func Int32ToUint32(v int32) (uint32, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int32(%d) to uint32: %w", v, ErrOverflowTypeConversion)
	}
	return uint32(v), nil
}

func Int32ToUint64(v int64) (uint64, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int32(%d) to uint64: %w", v, ErrOverflowTypeConversion)
	}
	return uint64(v), nil
}

func Int16ToInt8(v int16) (int8, error) {
	if v < -1*(1<<7) || (1<<7) <= v {
		return 0, fmt.Errorf("failed to convert int16(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

func Int16ToUint8(v int16) (uint8, error) {
	if v < 0 || (1<<8) <= v {
		return 0, fmt.Errorf("failed to convert int16(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Int16ToUint16(v int16) (uint16, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int16(%d) to uint16: %w", v, ErrOverflowTypeConversion)
	}
	return uint16(v), nil
}

func Int16ToUint32(v int16) (uint32, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int16(%d) to uint32: %w", v, ErrOverflowTypeConversion)
	}
	return uint32(v), nil
}

func Int16ToUint64(v int16) (uint64, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int16(%d) to uint64: %w", v, ErrOverflowTypeConversion)
	}
	return uint64(v), nil
}

func Int8ToUint8(v int8) (uint8, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int8(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Int8ToUint16(v int8) (uint16, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int8(%d) to uint16: %w", v, ErrOverflowTypeConversion)
	}
	return uint16(v), nil
}

func Int8ToUint32(v int8) (uint32, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int8(%d) to uint32: %w", v, ErrOverflowTypeConversion)
	}
	return uint32(v), nil
}

func Int8ToUint64(v int8) (uint64, error) {
	if v < 0 {
		return 0, fmt.Errorf("failed to convert int8(%d) to uint64: %w", v, ErrOverflowTypeConversion)
	}
	return uint64(v), nil
}

func Uint64ToInt8(v uint64) (int8, error) {
	if (1 << 7) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

func Uint64ToInt16(v uint64) (int16, error) {
	if (1 << 15) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to int16: %w", v, ErrOverflowTypeConversion)
	}
	return int16(v), nil
}

func Uint64ToInt32(v uint64) (int32, error) {
	if (1 << 31) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to int32: %w", v, ErrOverflowTypeConversion)
	}
	return int32(v), nil
}

func Uint64ToInt64(v uint64) (int64, error) {
	if (1 << 63) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to int64: %w", v, ErrOverflowTypeConversion)
	}
	return int64(v), nil
}

func Uint64ToUint8(v uint64) (uint8, error) {
	if (1 << 8) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Uint64ToUint16(v uint64) (uint16, error) {
	if (1 << 16) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to uint16: %w", v, ErrOverflowTypeConversion)
	}
	return uint16(v), nil
}

func Uint64ToUint32(v uint64) (uint32, error) {
	if (1 << 32) <= v {
		return 0, fmt.Errorf("failed to convert uint64(%d) to uint32: %w", v, ErrOverflowTypeConversion)
	}
	return uint32(v), nil
}

func Uint32ToInt8(v uint32) (int8, error) {
	if (1 << 7) <= v {
		return 0, fmt.Errorf("failed to convert uint32(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

func Uint32ToInt16(v uint32) (int16, error) {
	if (1 << 15) <= v {
		return 0, fmt.Errorf("failed to convert uint32(%d) to int16: %w", v, ErrOverflowTypeConversion)
	}
	return int16(v), nil
}

func Uint32ToInt32(v uint64) (int32, error) {
	if (1 << 31) <= v {
		return 0, fmt.Errorf("failed to convert uint32(%d) to int32: %w", v, ErrOverflowTypeConversion)
	}
	return int32(v), nil
}

func Uint32ToUint8(v uint32) (uint8, error) {
	if (1 << 8) <= v {
		return 0, fmt.Errorf("failed to convert uint32(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Uint32ToUint16(v uint32) (uint16, error) {
	if (1 << 16) <= v {
		return 0, fmt.Errorf("failed to convert uint32(%d) to uint16: %w", v, ErrOverflowTypeConversion)
	}
	return uint16(v), nil
}

func Uint16ToInt8(v uint16) (int8, error) {
	if (1 << 7) <= v {
		return 0, fmt.Errorf("failed to convert uint16(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

func Uint16ToInt16(v uint16) (int16, error) {
	if (1 << 15) <= v {
		return 0, fmt.Errorf("failed to convert uint16(%d) to int16: %w", v, ErrOverflowTypeConversion)
	}
	return int16(v), nil
}

func Uint16ToUint8(v uint16) (uint8, error) {
	if (1 << 8) <= v {
		return 0, fmt.Errorf("failed to convert uint16(%d) to uint8: %w", v, ErrOverflowTypeConversion)
	}
	return uint8(v), nil
}

func Uint8ToInt8(v uint8) (int8, error) {
	if (1 << 7) <= v {
		return 0, fmt.Errorf("failed to convert uint16(%d) to int8: %w", v, ErrOverflowTypeConversion)
	}
	return int8(v), nil
}

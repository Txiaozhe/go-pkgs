package encoding

const (
	encodedNull = 0x00
	encodedNotNull = 0x01

	floatNaN = encodedNotNull + 1
	floatNeg = floatNaN + 1
	floatZero = floatNeg + 1
	floatPos = floatZero + 1
	floatNaNDesc = floatPos + 1
)

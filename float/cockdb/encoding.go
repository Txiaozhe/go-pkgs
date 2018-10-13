package encoding

const (
	encodedNull    = 0x00
	encodedNotNull = 0x01

	floatNaN     = encodedNotNull + 1
	floatNeg     = floatNaN + 1
	floatZero    = floatNeg + 1
	floatPos     = floatZero + 1
	floatNaNDesc = floatPos + 1

	encodedNullDesc = 0xff
)

type Type int

const (
	Unknown   Type = 0
	Null      Type = 1
	NotNull   Type = 2
	Int       Type = 3
	Float     Type = 4
	Decimal   Type = 5
	Bytes     Type = 6
	BytesDesc Type = 7 // Bytes encoded descendingly
	Time      Type = 8
	Duration  Type = 9
	True      Type = 10
	False     Type = 11
	UUID      Type = 12
	Array     Type = 13
	IPAddr    Type = 14
	// SentinelType is used for bit manipulation to check if the encoded type
	// value requires more than 4 bits, and thus will be encoded in two bytes. It
	// is not used as a type value, and thus intentionally overlaps with the
	// subsequent type value. The 'Type' annotation is intentionally omitted here.
	SentinelType      = 15
	JSON         Type = 15
	Tuple        Type = 16
	BitArray     Type = 17
	BitArrayDesc Type = 18 // BitArray encoded descendingly
)

func EncodeUint64Ascending(b []byte, v uint64) []byte {
	return append(b,
		byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32),
		byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func DecodeUint64Ascending(b []byte) ([]byte, uint64, error) {
	leftover, v, err := DecodeUint64Ascending(b)
	return leftover, ^v, err
}

func PeekType(b []byte) Type {
	if len(b) >= 1 {
		m := b[0]
		switch {
		case m == encodedNull, m == encodedNullDesc:
			return Null

			// TODO:
		}
	}
	return Unknown
}

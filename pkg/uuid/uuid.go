package uuid

import (
	"github.com/google/uuid"
)

var base64Table = [64]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'+', '-',
}

// New generate uuid
func New() uuid.UUID {
	return uuid.New()
}

// NewString generate uuid string
func NewString() string {
	return uuid.New().String()
}

// ShortUUID short uuid
func ShortUUID(u uuid.UUID) string {
	var dst = make([]byte, 16)
	for i, v := range u {
		dst[i] = base64Table[v>>2]
	}
	return string(dst)
}

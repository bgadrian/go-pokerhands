package five

import (
	"fmt"
	"hash/crc32"
)

var crc32q = crc32.MakeTable(0xD5828281)

// FiveHandHashFromStructs TODO needs to be optimised, use the string directly from the API
func FiveHandHashFromStructs(h FiveHand) uint32 {
	asString := fmt.Sprintf("%s%s%s%s%s%s", h[0], h[1], h[2], h[3], h[4])
	return crc32.Checksum([]byte(asString), crc32q)
}

package internal

import (
	"encoding/base64"
	"github.com/spaolacci/murmur3"
)

// todo : implement some collission testing strategy like bloom filter and salt for similar urls
func generateHashValue(url string) string {
    hash := murmur3.Sum64([]byte(url))
    shortHash := hash & 0x3FFFFFFFFFF // Mask to get the first 42 bits (7 Base64 chars)

    buf := make([]byte, 8) 
    for i := 0; i < 6; i++ {
        buf[i] = byte(shortHash >> (8 * (5 - i)))
    }

    return base64.URLEncoding.EncodeToString(buf)[:7]
}
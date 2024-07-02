package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func POW_sha256_hash_example() {
	nickname := "Taylor Li"
	// first hash starts with 0000
	findRepeatZeroHash(nickname, 4)
	// first hash starts with 00000
	findRepeatZeroHash(nickname, 5)
}

func findRepeatZeroHash(
	nickname string,
	repeatCout int,
) {
	nonce := 0
	startTime := time.Now()
	zeroToValidate := strings.Repeat("0", repeatCout)
	for nonce = 0; ; nonce++ {
		hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", nickname, nonce)))
		hashHex := hex.EncodeToString(hash[:])
		fmt.Printf("%x\r", hash)
		if strings.HasPrefix(hashHex, zeroToValidate) {
			fmt.Printf("First '%s' hash: %x\n", zeroToValidate, hash)
			fmt.Printf("Nonce: %d\n", nonce)
			fmt.Printf("Time taken: %v\n", time.Since(startTime))
			break
		}
	}
}

package main

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"
)

func MakeRandomString() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

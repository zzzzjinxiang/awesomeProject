package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func mainxo() {
	for i := 0; i < 1000000000000000; i++ {
		data := sha256.Sum256([]byte(strconv.Itoa(i)))
		fmt.Printf("%10d, %x\n", i, data)
	}
}

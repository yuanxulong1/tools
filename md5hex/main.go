package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	arg_num := len(os.Args)
	if arg_num < 2 {
		fmt.Println("usage:md5hex string")
		os.Exit(1)
	}
	h := md5.New()
	h.Write([]byte(os.Args[1]))
	cipherStr := h.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr))
}

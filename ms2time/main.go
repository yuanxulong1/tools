package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	arg_num := len(os.Args)
	if arg_num < 2 {
		fmt.Println("usage:ms2time timeinmillseconds")
		os.Exit(1)
	}
	timeinms, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("参数出错")
		os.Exit(1)
	}
	t := time.Unix(int64(timeinms/1000), int64(timeinms%(1000*1000)))
	fmt.Println(t)
}

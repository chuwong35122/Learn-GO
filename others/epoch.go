package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	fmt.Println(sec)

}

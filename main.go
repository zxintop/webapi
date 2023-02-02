package main

import (
	"fmt"
	"github.com/zxintop/gocommon/strutil"
)

func main() {
	str := strutil.ToUpper("Hello, world")
	fmt.Println(str)

	str = strutil.ToLower(str)
	fmt.Println(str)
}

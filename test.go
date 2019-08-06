package main

import (
	"fmt"
	"strings"
)

func main()  {
	t := strings.Split("tools/accounts", "/")
	fmt.Println(t[0])
	fmt.Println(t[1])
}

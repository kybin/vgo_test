package main

import (
	"fmt"

	"github.com/kybin/blackfriday/v2"
)

func main() {
	html := blackfriday.Run([]byte("#Title"))
	fmt.Println(html)
}

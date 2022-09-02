package main

import (
	"fmt"

	"github.com/alibabacloud-go/darabonba-openapi/client"
)

func main() {
	c, _ := client.NewClient(nil)
	fmt.Println(c.Endpoint)

	a := "abcdefg"
	fmt.Println(a[:len(a)-2])
}

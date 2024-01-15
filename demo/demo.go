package main

import (
	"fmt"
	"github.com/NubeIO/rxcleint"
)

func main() {
	c, err := rxcleint.New()
	fmt.Println(err)
	resp := c.IPValidation("162")
	fmt.Printf("%+v\n", resp)
}

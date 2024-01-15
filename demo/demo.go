package main

import (
	"fmt"
	"github.com/NubeIO/rxclient"
)

func main() {
	c, err := rxclient.New()
	fmt.Println(err)
	resp := c.IPValidation("162")
	fmt.Printf("%+v\n", resp)
}

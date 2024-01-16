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
	command, err := c.RunCommand(&rxclient.CommandBody{
		Command: "pwd",
		Args:    nil,
		Timeout: 0,
	})
	fmt.Println(err)
	fmt.Printf("%+v\n", command)

	cmdResp, err := c.SystemdCommand("mosquitto", rxclient.SystemCTLStart, 1)
	fmt.Println(err)
	fmt.Printf("%+v\n", cmdResp)

	statusResp, err := c.SystemdStatus("mosquitto", 1)
	fmt.Println(err)
	fmt.Printf("%+v\n", statusResp)

}

package rxclient

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := New()
	fmt.Println(err)
	//
	//resp := c.IPValidation("123423")
	//fmt.Printf("%+v\n", resp)

	c.UserAll()
}

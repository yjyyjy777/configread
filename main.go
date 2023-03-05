package main

import (
	"configread/parsglobal"
	"fmt"
)

func main() {
	c := parsglobal.Globalconfig()
	fmt.Println(c)
}

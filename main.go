package main

import (
  "flag"
	"fmt"
//  "os"
)

func init() {
  var add = flag.NewFlagSet("add", flag.ExitOnError)
}

func logo() {
	fmt.Println(" ,                  ")
	fmt.Println("-+- _ ._ _ ._ . . __")
	fmt.Println(" | (/,[ | )[_)(_|_) ")
	fmt.Println("           |        ")

	fmt.Println("		~0xbiel")
}

func main() {
	logo()
 // init()
}

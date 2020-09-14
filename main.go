package main

import (
	"fmt"
  "os"
)

func logo() {
	fmt.Println(" ,                  ")
	fmt.Println("-+- _ ._ _ ._ . . __")
	fmt.Println(" | (/,[ | )[_)(_|_) ")
	fmt.Println("           |        ")

	fmt.Println("		~0xbiel\n")
}

func help() {
  logo()
  fmt.Println("[-a] add new entry.\n")
  fmt.Println("[-i] init time.\n")
  fmt.Println("[-s] stop time.\n")
  fmt.Println("[-r] remove entry.\n")
  fmt.Println("[-l] list entries.\n")
  fmt.Println("[-e] export activity.\n")
  fmt.Println("[-h] show help.\n")
}

func main() {
  if (len(os.Args) < 2) {
    help() 
  }
}


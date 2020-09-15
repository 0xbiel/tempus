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
  
  fmt.Println("usage: tempus [args]\n\n")
  fmt.Println("[-a] [add] - add new entry.\n")
  fmt.Println("[-i] [init] - init time.\n")
  fmt.Println("[-s] [stop] - stop time.\n")
  fmt.Println("[-r] [remove] - remove entry.\n")
  fmt.Println("[-l] [list] - list entries.\n")
  fmt.Println("[-e] [export] - export activity.\n")
  fmt.Println("[-h] [help] - show help.\n")
}

func main() {

  if (len(os.Args) < 2) {

    help()

  } else if (os.Args[1] == "-a" || os.Args[1] == "add") {

      fmt.Println("add\n");

  }
}


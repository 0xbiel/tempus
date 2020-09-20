package main

import (
	"fmt"
  "os"
)

func logo() {
  logo := `
 ,                  
-+- _ ._ _ ._ . . __
 | (/,[ | )[_)(_|_) 
           |        
		~0xbiel
  `

	fmt.Println(logo)
}

func help() {
  logo()
  fmt.Println("usage: tempus [args]\n\n")
  fmt.Println("[-a] - add new entry.\n")
  fmt.Println("[-i] - init time.\n")
  fmt.Println("[-s] - stop time.\n")
  fmt.Println("[-r] - remove entry.\n")
  fmt.Println("[-l] - list entries.\n")
  fmt.Println("[-e] - export activity.\n")
  fmt.Println("[-h] - show help.\n")
}

func main() {
  if (len(os.Args) < 2) {

    help();

  } else {

      switch os.Args[1] {
        case "-a": 
          fmt.Println("add")
        case "-i":
          fmt.Println("init")
        case "-s":
          fmt.Println("stop")
        case "-r":
          fmt.Println("remove")
        case "-l":
          fmt.Println("list")
        case "-e":
          fmt.Println("export")
        case "-h":
          help()
        default:
          help()
      }
    }
}

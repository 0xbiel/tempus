package main

import (
//  add "./add"
	"fmt"
  "os"
  tui "./tui"
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
  fmt.Println("Usage: tempus [args]\n")

  fmt.Println("CLI:")
  fmt.Println("----\n")

  fmt.Println("   [-h] - show help.\n")

  fmt.Println("TUI:")
  fmt.Println("----\n")

  fmt.Println("   [a] - add new entry.\n")
  fmt.Println("   [i] - init time.\n")
  fmt.Println("   [s] - stop time.\n")
  fmt.Println("   [r] - remove entry.\n")
  fmt.Println("   [l] - list entries.\n")
  fmt.Println("   [e] - export activity.\n")
}

func main() {
  if (len(os.Args) < 2) {

    tui.Start(); 

  } else if (os.Args[1] != "-h") {

      fmt.Println("argument not recognized.")

    } else {

      help()

    }
}

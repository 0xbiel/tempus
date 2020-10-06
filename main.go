package main

import (
//  i "./init"
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
  fmt.Println("Usage: tempus [args]\n\n")

  fmt.Println("CLI")
  fmt.Println("---")
  fmt.Println("   [-h] - show help.\n")

  fmt.Println("TUI")
  fmt.Println("---")

  fmt.Println("   [a] - add new entry.")
  fmt.Println("   [i] - init time.")
  fmt.Println("   [s] - stop time.")
  fmt.Println("   [r] - remove entry.")
  fmt.Println("   [l] - list entries.")
  fmt.Println("   [e] - export activity.")
}

func main() {
  if (len(os.Args) < 2) {

    tui.Start(); 

  } else if(os.Args[1] != "-h"){

      fmt.Println("argument not recognized.")

  } else {

      help()

  }
}

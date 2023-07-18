package main

import (
	"flag"
	"fmt"

	tools "github.com/claybex/go-cli/tools"
)

var mod = flag.Int("mod",0,"select relevant working mode")
var help = flag.String("help","","help with app usage")

func main() {
	flag.Parse()
	if *help != "" {
		fmt.Print("Usage: --mod {1,2,3,4,5}\n 1 - TCP port scanner \n 2 - tcp proxy\n" + 
		"3 - Http Client \n")	
	} else {

	switch *mod {
	case 1:
		var choise1 string
		var choise2 string
		fmt.Println("Welcome to tcp scanner!")
		fmt.Println("Enter address and port in format: 8.8.8.8:80 or google.com:443")
		fmt.Scanf("%s %s",&choise1, &choise2)
		tools.TcpScanner(choise1, choise2)
		
	case 2:
		fmt.Println("Welcome to tcp !")
	case 3:
		fmt.Println("Welcome to http client!")
	default: 
		fmt.Println("Invalid option specified! --help for usage")
	}
	}
}




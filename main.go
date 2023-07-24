package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	tools "github.com/claybex/go-cli/tools"
)

var ( 
mod = flag.Int("mod",0,"select relevant working mode")
)

func main() {
	flag.Parse()
	if *mod == 0 {
		fmt.Print("Usage: --mod {1,2,3,4,5}\n1 - TCP port scanner \n2 - tcp proxy\n" + 
		"3 - Http Client \n")	
	} else {
	switch *mod {
	case 1:
		// cli-tools --mod 1 google.com 80
		
		if len(os.Args) == 4 {
			arg1 := os.Args[3]
			fmt.Printf("Scanning ports 0-65535 for %s\n", arg1)
		    tools.TCPScanner(arg1, 0, 65535)
		} else if len(os.Args) == 5 {
			if arg1, arg2 := os.Args[3], os.Args[4]; arg1 != "" && arg2 != "" {
				port, _ := strconv.Atoi(arg2) 
				tools.SingleTcpScanner(arg1, port)
			}
			// cli-tools --mod 1 google.com 0 1024 
		} else if len(os.Args) == 6 {
			if arg1, arg2, arg3 := os.Args[3], os.Args[4], os.Args[5]; arg1 != "" && arg2 < arg3{
				port1, _ := strconv.Atoi(arg2)
				port2, _ := strconv.Atoi(arg3)
				tools.TCPScanner(arg1, port1, port2)
			} else {
				log.Fatalln("Error, invalid arguments!")
			}
		} else {
				var (choise1 string
					x,y int
				)
				fmt.Println("Welcome to tcp scanner!")
				fmt.Println("Enter address and port in format: 8.8.8.8:80 or google.com:443")
				fmt.Scanf("%s %d %d",&choise1, &x, &y)
				if x != 0 && y == 0 {
					tools.SingleTcpScanner(choise1, x)
				} else if x == 0 && y == 0 {
					y = 65535
				}
			tools.TCPScanner(choise1, x,  y )
			}

	case 2:
		fmt.Println("Welcome to tcp proxy!")
		
	case 3:
		fmt.Println("Welcome to http client!")
	default: 
		fmt.Println("Invalid option specified! --help for usage")
	}
	}
	// var wg sync.WaitGroup
	// 	wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Print("Game over!")
	// }()
	// wg.Wait()

}




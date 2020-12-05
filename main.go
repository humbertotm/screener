package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// TODO: Incorporate logging utility instead of just printing lines
	fmt.Println("Starting screener..")

	coDetailsCmd := flag.NewFlagSet("companydetails", flag.ExitOnError)
	coCIK := coDetailsCmd.String("cik", "", "cik")
	coDetailsCmd.Usage = func() {
		fmt.Println("cik flag not provided")
		fmt.Println("Usage: -cik string")
	}

	if len(os.Args) < 2 {
		fmt.Println("Expected fullscreen or companydetails subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	// go run main.go fullscreen
	case "fullscreen":
		fmt.Println("Executing full screening task...")
	// go run main.go companydetails --cik=111111
	case "companydetails":
		if len(os.Args[2:]) < 1 {
			coDetailsCmd.Usage()

		} else {
			coDetailsCmd.Parse(os.Args[2:])
			fmt.Printf("Extracting details for CIK %s\n", *coCIK)
		}

	default:
		fmt.Println("Expected fullscreen or companydetails subcommands")
		os.Exit(1)
	}

	os.Exit(0)
}

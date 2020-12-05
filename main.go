package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// TODO: Incorporate logging utility instead of just printing lines
	fmt.Println("Starting screener..")

	// Initiate context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initiate MongoDB connection
	client, err := initMongoClient(ctx)
	if err != nil {
		fmt.Printf("Error while establishing connection with MongoDB: %s\n", err.Error())
		os.Exit(1)
	}

	// Terminate MongoDB connection
	// TODO: gracefully handle this panic?
	defer func() {
		fmt.Println("Terminating mongodb connection")
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Handle input subcommand
	handleCommand(client)

	// Exit program successfully
	os.Exit(0)
}

func initMongoClient(ctx context.Context) (*mongo.Client, error) {
	clientOpts := options.Client()
	clientOpts.ApplyURI("mongodb://localhost:27017")
	clientOpts.SetAuth(options.Credential{
		Username: "mongoadmin",
		Password: "mongoadmin",
	})
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// TODO: return errors and let main function handle the exit
func handleCommand(client *mongo.Client) {
	coDetailsCmd := flag.NewFlagSet("companydetails", flag.ExitOnError)
	coCIK := coDetailsCmd.String("cik", "", "cik")
	coDetailsCmd.Usage = func() {
		fmt.Println("cik flag not provided")
		fmt.Println("Usage: -cik string")
	}

	if len(os.Args) < 2 {
		// A more apprpriate Usage handling function call would be better here
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
			os.Exit(1)

		} else {
			coDetailsCmd.Parse(os.Args[2:])
			fmt.Printf("Extracting details for CIK %s\n", *coCIK)
		}

	default:
		fmt.Println("Expected fullscreen or companydetails subcommands")
		os.Exit(1)
	}
	return
}

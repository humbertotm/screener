package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"screener.com/profile/repository/mongodb"
	"screener.com/screener/delivery"
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
		// CURIOSITY: this line is printed whenever there's a panic in the
		// handleCommand function, not when the program exits successfully.
		// Why? Is this defered function being executed upon successful termination?
		fmt.Println("Terminating mongodb connection")
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Handle input subcommand
	if err := handleCommand(ctx, client); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

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
func handleCommand(ctx context.Context, client *mongo.Client) error {
	coDetailsCmd := flag.NewFlagSet("companydetails", flag.ExitOnError)
	coCIK := coDetailsCmd.String("cik", "", "cik")
	coDetailsCmd.Usage = func() {
		fmt.Println("cik flag not provided")
		fmt.Println("Usage: -cik string")
	}

	if len(os.Args) < 2 {
		// A more apprpriate Usage handling function call would be better here
		return fmt.Errorf("Expected fullscreen or companydetails subcommands")
	}

	switch os.Args[1] {
	// go run main.go fullscreen
	case "fullscreen":
		fmt.Println("Executing full screening task...")
		r := mongodb.NewProfileRepository(client)
		_, err := r.GetFullCIKList(ctx)
		if err != nil {
			return err
		}
	// go run main.go companydetails --cik=111111
	case "companydetails":
		if len(os.Args[2:]) < 1 {
			coDetailsCmd.Usage()
			return fmt.Errorf("Invalid arguments")

		} else {
			coDetailsCmd.Parse(os.Args[2:])
			fmt.Printf("Extracting details for CIK %s\n", *coCIK)
			handler := delivery.NewScreenerHandler(client)
			if err := handler.GetStatsForCIK(ctx, *coCIK); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("Expected fullscreen or companydetails subcommands")
	}

	return nil
}

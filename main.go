package main

import (
	"context"
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

//Color define color
type Color string

//const
const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {

	projectID := flag.StringP("project", "p", "", "Select project where your secret lives")
	SecretName := flag.StringP("secretname", "s", "", "Secret name you wish to retreive")

	flag.Parse()

	GoogleCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	if len(GoogleCreds) == 0 {
		colorize(ColorRed, "ERR: add to environment GOOGLE_APPLICATION_CREDENTIALS variable with your service account credentials")
		os.Exit(2)
	}
	if len(*SecretName) == 0 {

		colorize(ColorRed, "ERR: Provide secret name")
		os.Exit(2)
	}

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", *projectID, *SecretName),
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret latest version: %v", err)
	}

	fmt.Printf("%s", result.Payload.Data)
}

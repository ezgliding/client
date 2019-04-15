package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// parseCmd represents the 'view' command.
var parseCmd = &cobra.Command{
	Use:   "parse [FILE]",
	Short: "parse the given flight",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Sets your Google Cloud Platform project ID.
		projectID := "YOUR_PROJECT_ID"

		// Creates a client.
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		// Sets the name for the new bucket.
		bucketName := "my-new-bucket"

		// Creates a Bucket instance.
		bucket := client.Bucket(bucketName)

		// Creates the new bucket.
		if err := bucket.Create(ctx, projectID, nil); err != nil {
			log.Fatalf("Failed to create bucket: %v", err)
		}

		fmt.Printf("Bucket %v created.\n", bucketName)
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
}

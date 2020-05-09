package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-clix/cli"
)

func main() {
	rootCmd := &cli.Command{
		Use:   "farto",
		Short: "Farto",
	}
	rootCmd.AddCommand(
		generateCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func generateCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "generate",
		Short: "Generate static site locally.",
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)
		if err != nil {
			return err
		}

		svc := s3.New(sess)

		b, err := walkBucket(svc, "farto.cloud")
		if err != nil {
			return err
		}

		fmt.Println(b)

		return nil
	}
	return cmd
}

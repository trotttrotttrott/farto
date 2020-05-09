package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/go-clix/cli"
	"gopkg.in/yaml.v2"
)

type config struct {
	S3Region string `yaml:"s3Region"`
	S3Bucket string `yaml:"s3Bucket"`
	S3Prefix string `yaml:"s3Prefix"`
}

func main() {
	rootCmd := &cli.Command{
		Use:   "farto",
		Short: "Farto",
	}
	rootCmd.AddCommand(
		siteCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func siteCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "site",
		Short: "Commands related to the static site.",
	}
	cmd.AddCommand(
		generateCmd(),
	)
	return cmd
}

func getConfig() (c config, err error) {
	y, err := ioutil.ReadFile("farto.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(y, &c)
	return
}

func generateCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "generate",
		Short: "Generate static site locally.",
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		c, err := getConfig()
		if err != nil {
			return err
		}
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(c.S3Region)},
		)
		if err != nil {
			return err
		}
		svc := s3.New(sess)
		b, err := walkBucket(svc, c.S3Bucket, c.S3Prefix)
		if err != nil {
			return err
		}

		fmt.Println(b)

		return nil
	}
	return cmd
}

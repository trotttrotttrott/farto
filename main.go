package main

import (
	"fmt"
	"log"

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
		b, err := walkBucket("farto.cloud")
		if err != nil {
			return err
		}
		fmt.Println(b)
		return nil
	}
	return cmd
}

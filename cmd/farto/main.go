package main

import (
	"log"

	"github.com/go-clix/cli"
	"github.com/trotttrotttrott/farto/pkg/farto"
)

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
		siteGenerateCmd(),
	)
	return cmd
}

func siteGenerateCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "generate",
		Short: "Generate static site locally.",
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		return farto.SiteGenerate()
	}
	return cmd
}

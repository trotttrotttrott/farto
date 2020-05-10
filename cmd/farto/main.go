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
		fartosCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

// site commands

func siteCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "site",
		Short: "Commands related to the static site.",
	}
	cmd.AddCommand(
		siteGenerateCmd(),
		sitePublishCmd(),
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

func sitePublishCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "publish",
		Short: "Publish static site to S3.",
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		return farto.SitePublish()
	}
	return cmd
}

// fartos commands

func fartosCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "fartos",
		Short: "Commands related to farto management.",
	}
	cmd.AddCommand(
		fartosNormalizeCmd(),
	)
	return cmd
}

func fartosNormalizeCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "normalize",
		Short: "Create normalized versions of your fartos.",
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		p := args[0]
		return farto.FartosNormalize(p)
	}
	return cmd
}

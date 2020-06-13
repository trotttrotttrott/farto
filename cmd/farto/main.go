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
		Args:  cli.ArgsExact(0),
	}
	customTemplate := cmd.Flags().StringP("custom-template", "t", "", "Path to custom template.")
	cmd.Run = func(cmd *cli.Command, args []string) error {
		return farto.SiteGenerate(customTemplate)
	}
	return cmd
}

func sitePublishCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "publish",
		Short: "Publish static site to S3.",
		Args:  cli.ArgsExact(0),
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
		fartosUploadCmd(),
	)
	return cmd
}

func fartosNormalizeCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "normalize <path>",
		Short: "Create normalized versions of your fartos.",
		Args:  cli.ArgsExact(1),
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		p := args[0]
		return farto.FartosNormalize(p)
	}
	return cmd
}

func fartosUploadCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "upload <path>",
		Short: "Upload original and normalized photos to S3.",
		Args:  cli.ArgsExact(1),
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		p := args[0]
		return farto.FartosUpload(p)
	}
	return cmd
}

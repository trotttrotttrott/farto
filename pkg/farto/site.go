package farto

import (
	"html/template"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type site struct {
	Title  string
	Fartos []string
}

func SiteGenerate() error {
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
	s := site{
		Title:  "Farto",
		Fartos: b,
	}
	tmpl, err := template.ParseFiles("pkg/farto/templates/index.html")
	if err != nil {
		return err
	}
	err = os.MkdirAll("site", 0755)
	if err != nil {
		return err
	}
	f, err := os.Create("site/index.html")
	if err != nil {
		return err
	}
	err = tmpl.Execute(f, s)
	return err
}

func SitePublish() error {
	return nil
}

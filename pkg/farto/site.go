package farto

import (
	"html/template"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type site struct {
	Title  string
	Fartos map[string][]string
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
	keys, err := walkBucket(svc, c.S3Bucket, c.S3Prefix)
	if err != nil {
		return err
	}

	fartos := map[string][]string{}
	for _, key := range keys {
		d, f := path.Split(key)
		if d != "/site/" && path.Ext(f) != "" {
			fartos[d] = append(fartos[d], key)
		}
	}

	s := site{
		Title:  "Farto",
		Fartos: fartos,
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
	upload(svc, c.S3Bucket, c.S3Prefix, "site")
	return nil
}

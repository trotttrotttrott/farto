package farto

import (
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type site struct {
	Title    string
	Headline string
	Copy     string
	Fartos   map[string][]string
}

const siteTmpl = `<!doctype html>
<html>
  <head>
    <title>{{.Title}}</title>
  </head>
  <body>
    <h1>{{.Headline}}</h1>
    <p>{{.Copy}}</p>
    {{range $folder, $paths := .Fartos}}
    <div style="padding: 20px 0;">
      <h4 style="text-align: left; margin-left: 50px;">{{$folder}}</h4>
      {{range $paths}}
      <a href="/{{$folder}}.farto.n/{{.}}.jpg" target="_blank">
        <img src="/{{$folder}}.farto.n.t/{{.}}.jpg" />
      </a>
      {{end}}
    </div>
    {{end}}
  </body>
</html>`

func SiteGenerate(customTemplatePath *string) error {
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
		d = strings.Trim(d, "/")
		fartos[d] = append(fartos[d], f)
	}
	s := site{
		Title:    c.SiteTitle,
		Headline: c.SiteHeadline,
		Copy:     c.SiteCopy,
		Fartos:   fartos,
	}
	var tmpl *template.Template
	if *customTemplatePath == "" {
		tmpl, err = template.New("index.html").Parse(siteTmpl)
	} else {
		tmpl, err = template.ParseFiles(*customTemplatePath)
	}
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

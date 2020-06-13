package farto

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/disintegration/imaging"
	"github.com/jdeng/goheif"
	"github.com/rwcarlsen/goexif/exif"
)

func FartosNormalize(p string) error {

	versions := map[int]string{
		800: "n",   // normalized
		200: "n.t", // " thumbnail
	}
	for size, dirSuffix := range versions {
		dir := fmt.Sprintf("%s.farto.%s", path.Clean(p), dirSuffix)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		versions[size] = dir
	}

	err := filepath.Walk(p, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			var src image.Image
			if strings.ToLower(path.Ext(info.Name())) == ".heic" {
				f, err := os.Open(p)
				if err != nil {
					return err
				}
				src, err = goheif.Decode(f)
				if err != nil {
					fmt.Printf("WARNING: There was a problem decoding %s...\n", p)
					fmt.Printf("WARNING: > %s.\n", err)
					fmt.Println("WARNING: You'll have to deal with this one manually.")
					return nil
				}
				b, err := goheif.ExtractExif(f)
				if err != nil {
					return err
				}
				r := bytes.NewReader(b)
				x, err := exif.Decode(r)
				if err != nil {
					return err
				}
				o, err := x.Get(exif.Orientation)
				if err != nil {
					return err
				}
				oi, err := o.Int(0)
				if err != nil {
					return err
				}
				switch oi {
				case 2:
					src = imaging.FlipH(src)
				case 4:
					src = imaging.FlipV(src)
				case 8:
					src = imaging.Rotate90(src)
				case 3:
					src = imaging.Rotate180(src)
				case 6:
					src = imaging.Rotate270(src)
				case 5:
					src = imaging.Transpose(src)
				case 7:
					src = imaging.Transverse(src)
				}
			} else {
				src, err = imaging.Open(p, imaging.AutoOrientation(true))
				if err != nil {
					return err
				}
			}
			for size, dir := range versions {
				img := imaging.Resize(src, 0, size, imaging.Lanczos)
				err = imaging.Save(img, path.Join(dir, fmt.Sprintf("%s.jpg", info.Name())))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func FartosUpload(p string) error {
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
	for _, dir := range []string{p, fmt.Sprintf("%s.farto.n", p), fmt.Sprintf("%s.farto.n.t", p)} {
		_, err := os.Stat(dir)
		if err != nil {
			return err
		}
		err = upload(svc, c.S3Bucket, c.S3Prefix, dir, false)
		if err != nil {
			return err
		}
	}
	return nil
}

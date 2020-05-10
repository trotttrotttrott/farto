package farto

import (
	"fmt"
	"image"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/jdeng/goheif"
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
					return err
				}
				src = imaging.Rotate270(src)
			} else {
				src, err = imaging.Open(p)
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

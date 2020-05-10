package farto

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
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
			src, err := imaging.Open(p)
			if err != nil {
				return err
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

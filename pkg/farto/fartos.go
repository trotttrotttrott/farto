package farto

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func FartosNormalize(p string) error {

	sizes := []int{800, 200}
	dirs := map[int]string{}
	for _, size := range sizes {
		dir := fmt.Sprintf("%s.farto.%d", path.Clean(p), size)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		dirs[size] = dir
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
			for _, size := range sizes {
				img := imaging.Resize(src, 0, size, imaging.Lanczos)
				err = imaging.Save(img, path.Join(dirs[size], fmt.Sprintf("%s.jpg", info.Name())))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

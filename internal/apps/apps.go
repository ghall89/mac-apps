package apps

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func GetApps() []fs.FileInfo {
	var res []fs.FileInfo

	usrDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	roots := []string{"/Applications", usrDir + "/Applications"}

	re := regexp.MustCompile(`\.app$`)

	for _, root := range roots {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			// If this directory is an app bundle, add it and don't descend into it
			if d.IsDir() && re.MatchString(d.Name()) {
				info, err := d.Info()
				if err != nil {
					return err
				}
				res = append(res, info)
				return filepath.SkipDir
			}

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	return res
}

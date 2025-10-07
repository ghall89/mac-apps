package apps

import (
	"bytes"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"howett.net/plist"
)

func GetApps() []AppResult {
	var res []AppResult

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

				name := info.Name()

				res = append(res, AppResult{Name: name, Path: root + `/` + name})
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

func GetInfo(path string) PlistRoot {
	filePath := path + "/Contents/Info.plist"

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(file)

	var data PlistRoot

	dec := plist.NewDecoder(r)
	if err := dec.Decode(&data); err != nil {
		log.Fatal(err)
	}

	return data
}

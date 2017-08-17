package smusic

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type mediaFolder struct {
	Path       string
	Songs      int
	MediaFiles []string
}

func (mf *mediaFolder) loadMediaFiles() error {
	// dh, err := os.Open(mf.Path)
	// defer dh.Close()
	// if err == nil {
	// 	fs, err := dh.Readdir(0)
	// 	if err == nil {
	// 		for _, fi := range fs {
	// 			if !fi.IsDir() && strings.HasSuffix(fi.Name(), ".mp3") {
	// 				mf.MediaFiles = append(mf.MediaFiles, fi.Name())
	// 			}
	// 		}
	// 	}
	// }
	err := filepath.Walk(mf.Path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
			mf.MediaFiles = append(mf.MediaFiles, path)
		}
		return nil
	})
	return err
}

func (mf *mediaFolder) getNextMediaFile() (mediaFileName string) {
	result := mf.MediaFiles[0]
	d, err := ioutil.ReadFile(mf.Path + "/.index")
	if err == nil {
		s := string(d)
		for i, f := range mf.MediaFiles {
			if f == s {
				if i < len(mf.MediaFiles)-1 {
					result = mf.MediaFiles[i+1]
					break
				}
			}
		}

	}
	ioutil.WriteFile(mf.Path+"/.index", []byte(result), 0644)
	return result
}

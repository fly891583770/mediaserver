package service

import (
	//"fmt"
	"io/ioutil"
	"path"

	"mediaserver/common/utils"
)

var (
	IMG_SUFFIXS   = []string{".jpg", ".bmp", ".jpeg", ".png", ".gif"}
	VIDEO_SUFFIXS = []string{".avi", ".rmvb", ".rm", ".asf", ".divx", ".mpg", ".mpeg", ".mpe", ".wmv", ".mp4", ".mkv", ".vob", ".h264"}
)

// return all images under basepath
func LsImages(basepath string) []string {

	return listFile(basepath, IMG_SUFFIXS)
}

// return all videos under basepath
func LsVideos(basepath string) []string {

	return listFile(basepath, VIDEO_SUFFIXS)
}

// recursive function, return a slice of full path of files under basepath and all its sub directory.
func listFile(basepath string, suffixs []string) (result []string) {

	files, _ := ioutil.ReadDir(basepath)

	for _, file := range files {

		//generate full path
		filepath := basepath + "/" + file.Name()

		if file.IsDir() {
			listFile(filepath, suffixs)
		} else {
			if suffixs == nil || utils.StrInSliceIgnoreCase(path.Ext(filepath), suffixs) {

				result = append(result, filepath)
			}
		}
	}

	return result
}

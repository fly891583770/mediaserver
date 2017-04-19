package service

import (
	"os"

	"mediaserver/entity"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//return a slice of Image, length of the slice
func genImages(imgpaths []string) (images []entity.Image, int) {

	for _, imgpath := range imgpaths {

		info, err := os.Stat(imgpath)
		if err != nil {
			logrus.Warnf("get FileInfo of [%s] failed, err: %v", imgpath, err)
			continue
		}
		image = entity.Image{ObjectId:bson.NewObjectId(), Name: info.Name(), Path: imgpath, ModTime: info.ModTime()}
		images = append(images, image)
	}

	return images, len(images)
}

package service

import (
	"fmt"
	"os"

	"mediaserver/common/utils"
	"mediaserver/entity"
)

// mongoDB page, db.test.find(xxx).sort({"num": 1}).skip(10).limit(10)
func ListImages(basepath string) (images []entity.Image, err error) {

}

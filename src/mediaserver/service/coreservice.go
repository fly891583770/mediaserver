package service

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	//"mediaserver/common/utils"
	"mediaserver/common/mongo"
	"mediaserver/entity"
)

const (
	IMAGE_COLLECTION_NAME = "images"
)

// mongoDB page, db.test.find(xxx).sort({"num": 1}).skip(10).limit(10)
func ListImages() ([]entity.Image, error) {

	documents := []entity.Image{}
	var selector = bson.M{}

	_, err := mongo.HandleQueryAll(&documents, mongo.QueryStruct{IMAGE_COLLECTION_NAME, selector, 0, 0, "_id"})
	if err != nil {
		logrus.Errorf("HandleQueryAll failed when list images, err: %v", err)
		return nil, err
	}

	return documents, nil

}

package mongo

import (
	"testing"

	"mediaserver/entity"

	"gopkg.in/mgo.v2/bson"
)

type TestStruct struct {
	ObjectId bson.ObjectId `bson:"_id" json:"_id"`
	Name     string        `bson:"name" json:"name"`
	Path     string        `bson:"path" json:"path"`
	Testkey  string        `bson:"testkey json:"testkey"`
}

func Init() {

	sessionMng := NewSessionManagerCustom("test", "../../mongo_config.toml")
	InjectOperator(sessionMng, "")
}

func TestHandleInsert(t *testing.T) {

	Init()
	document := &entity.Image{ObjectId: bson.NewObjectId(), Name: "hello2.jpg", Path: "/path/to/hello2.jpg"}
	//document := &TestStruct{ObjectId: bson.NewObjectId(), Name: "hello3.jpg", Path: "/path/to/hello3.jpg", Testkey: "ttttest"}
	HandleInsert("testcoll", document)
}
func TestHandleQueryOne(t *testing.T) {

	Init()
	document := &entity.Image{}
	var selector = bson.M{}
	selector[operator.ParamID()] = bson.ObjectIdHex("58f421c0e1382328c2bc7856")
	err := HandleQueryOne(&document, QueryStruct{"testcoll", selector, 0, 0, ""})
	t.Logf("document: %v", document)
	if err != nil {
		t.Errorf("failed, err: %v", err)
	}

}

func TestHandleQueryAll(t *testing.T) {

	Init()
	documents := []entity.Image{}
	var selector = bson.M{}
	selector["name"] = "hello2.jpg"
	total, err := HandleQueryAll(&documents, QueryStruct{"testcoll", selector, 0, 0, "_id"})
	t.Logf("total: %d, documents: %v", total, documents)
	if err != nil {
		t.Errorf("failed, err: %v", err)
	}

}

func TestHandleUpdateOne(t *testing.T) {
	Init()
	document := &entity.Image{ObjectId: bson.ObjectIdHex("58f421c0e1382328c2bc7856"), Name: "helloupdate.jpg", Path: "/path/to/helloupdate.jpg"}
	var selector = bson.M{}
	selector[operator.ParamID()] = bson.ObjectIdHex("58f421c0e1382328c2bc7856")
	created, err := HandleUpdateOne(&document, QueryStruct{"testcoll", selector, 0, 0, ""})
	t.Logf("create: %v", created)
	if err != nil {
		t.Errorf("failed, err: %v", err)
	}
}

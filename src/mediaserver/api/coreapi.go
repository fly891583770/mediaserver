package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mediaserver/entity"
	"mediaserver/service"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

const (
	HTTP_METHOD_GET    = "GET"
	HTTP_METHOD_POST   = "POST"
	HTTP_METHOD_PUT    = "PUT"
	HTTP_METHOD_DELETE = "DELETE"
)

type Resource struct {
}

func (r Resource) Register(router *mux.Router) {

	//router = router.Path("/mserver").Subrouter()
	// example: ip:port/images?start=0&count=10
	router.HandleFunc("/images", r.ListImgHandler).Methods(HTTP_METHOD_GET).Queries("start", "{start:[0-9]+}", "count", "{count:[0-9]+}")

}

func (r Resource) ListImgHandler(writer http.ResponseWriter, request *http.Request) {

	logrus.Infof("hello")

	//parse queries as a map into request.Form
	request.ParseForm()

	start, _ := strconv.Atoi(request.Form.Get("start"))
	count, _ := strconv.Atoi(request.Form.Get("count"))

	images, err := service.ListImages(start, count)
	if err != nil {
		logrus.Errorf("ListImages error, err: %v", err)
		writer.Write([]byte("internal err."))
		return
	}

	resp := entity.ImagesResp{Start: start, Count: count, Total: 100, Images: images}

	respbytes, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("marshal response error, err: %v", err)
		writer.Write([]byte("internal err."))
		return
	}

	writer.Write(respbytes)
}

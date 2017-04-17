package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mediaserver/entity"

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

	router.HandleFunc("/images", r.ListImgHandler).Methods(HTTP_METHOD_GET).Queries("start", "{start:[0-9]+}", "count", "{count:[0-9]+}")

}

func (r Resource) ListImgHandler(writer http.ResponseWriter, request *http.Request) {

	logrus.Infof("hello")

	//parse queries as a map into request.Form
	request.ParseForm()

	start, _ := strconv.Atoi(request.Form.Get("start"))
	count, _ := strconv.Atoi(request.Form.Get("count"))

	resp := entity.ImagesResp{Start: start, Count: count, Total: 100, Images: nil}

	respbytes, err := json.Marshal(resp)
	if err != nil {

		writer.Write([]byte("internal err."))
	}

	writer.Write(respbytes)
}

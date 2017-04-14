package main

import (
	//"log"
	//"net/http"
	"fmt"

	//"mediaserver/api"
	"mediaserver/service"

	//"github.com/gorilla/mux"
)

func main() {

	fmt.Println(service.LsImages("/home/mao/test"))

	//r := api.Resource{}
	//router := mux.NewRouter()
	//r.Register(router)

	//	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	//		t, err := route.GetPathTemplate()
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println(t)
	//		return nil
	//	})
	//log.Fatal(http.ListenAndServe(":18080", router))
}

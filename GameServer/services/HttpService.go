package services

import (
	"github.com/pzqf/zEngine/zObject"
)

type HttpService struct {
	zObject.BaseObject
}

func NewHttpService() *HttpService {
	a := &HttpService{}
	a.SetId(ServiceIdHttpServer)
	return a
}

func (ts *HttpService) Init() error {

	return nil
}

func (ts *HttpService) Close() error {

	return nil
}

func (ts *HttpService) Serve() {

}

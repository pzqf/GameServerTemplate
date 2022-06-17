package services

import (
	"github.com/pzqf/zEngine/zObject"
)

type RpcService struct {
	zObject.BaseObject
}

func NewRpcService() *RpcService {
	a := &RpcService{}
	a.SetId(ServiceIdRpcServer)
	return a
}

func (ts *RpcService) Init() error {

	return nil
}

func (ts *RpcService) Close() error {

	return nil
}

func (ts *RpcService) Serve() {

}

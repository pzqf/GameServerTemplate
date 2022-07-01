package services

import (
	"ServerTemplate/GameServer/config"
	"encoding/json"
	"fmt"
	"github.com/pzqf/zEngine/zLog"
	"github.com/pzqf/zEngine/zNet"
	"github.com/pzqf/zEngine/zObject"
	"go.uber.org/zap"
	"io"
	"net/http"
	"reflect"
	"runtime"
)

type HttpService struct {
	zObject.BaseObject
	httpServer *zNet.HttpServer
}

func NewHttpService() *HttpService {
	a := &HttpService{}
	a.SetId(ServiceIdHttpServer)
	return a
}

func (ts *HttpService) Init() error {
	if config.GConfig.HttpServer.Addr == "" {
		return nil
	}

	ts.httpServer = zNet.NewHttpServer(config.GConfig.HttpServer.Addr)

	for k, v := range zNet.GetHandler() {
		route := fmt.Sprintf("/%d", k)
		ts.httpServer.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			data, _ := io.ReadAll(request.Body)
			defer request.Body.Close()
			v(zNet.NewHttpSession(writer), k, data)
		})
		zLog.Info("http register", zap.String("route", route), zap.String("func", runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()))
	}

	//custom add
	ts.httpServer.HandleFunc("/config", func(writer http.ResponseWriter, request *http.Request) {
		b, _ := json.Marshal(config.GConfig)
		_, _ = writer.Write(b)
	})

	return nil
}

func (ts *HttpService) Close() error {
	if ts.httpServer != nil {
		ts.httpServer.Close()
	}
	return nil
}

func (ts *HttpService) Serve() {
	if ts.httpServer != nil {
		_ = ts.httpServer.Start()
	}

	zLog.Info("http info", zap.String("info", fmt.Sprintf("http server listing on %s", config.GConfig.HttpServer.Addr)))
}

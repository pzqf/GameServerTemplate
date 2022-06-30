package handler

import (
	"ServerTemplate/protocol"
	"encoding/json"
	"github.com/pzqf/zEngine/zNet"
)

func Init() error {
	zNet.InitDispatcherWorkerPool(10000)

	err := RegisterPlayerHandler()
	if err != nil {
		return err
	}

	err = zNet.RegisterHandler(protocol.TestPing, TestPing)
	if err != nil {
		return err
	}

	return nil
}

func TestPing(session zNet.Session, protoId int32, data []byte) {
	var reqData protocol.TestPingReq

	_ = json.Unmarshal(data, &reqData)
	resData := protocol.TestPingRes{
		Id:   reqData.Id,
		Name: reqData.Name,
		Time: reqData.Time,
	}
	d, _ := json.Marshal(resData)
	_ = session.Send(protocol.TestPing, d)
}

package handler

import (
	"ServerTemplate/protocol"
	"github.com/pzqf/zEngine/zNet"
)

func RegisterPlayerHandler() error {
	err := zNet.RegisterHandler(protocol.PlayerLogin, PlayerLogin)
	if err != nil {
		return err
	}
	err = zNet.RegisterHandler(protocol.PlayerLogout, PlayerLogout)
	if err != nil {
		return err
	}

	return nil
}

func PlayerLogin(session zNet.Session, protoId int32, data []byte) {

}

func PlayerLogout(session zNet.Session, protoId int32, data []byte) {

}
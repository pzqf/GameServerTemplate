package module

import (
	"ServerTemplate/GameServer/module/player"
	"github.com/pzqf/zEngine/zNet"
	"github.com/pzqf/zEngine/zObject"
	"sync"
)

type PlayerManager struct {
	zObject.ObjectManager
}

var GPlayerManager *PlayerManager
var once sync.Once

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{}
}

func GetPlayerManager() *PlayerManager {
	once.Do(func() {
		GPlayerManager = NewPlayerManager()
	})

	return GPlayerManager
}

func (pm *PlayerManager) AddPlayer(p *player.Player) error {
	err := pm.AddObject(p.GetId(), p)
	if err != nil {
		return err
	}
	return nil
}

func (pm *PlayerManager) GetPlayerById(id int) *player.Player {
	object, err := pm.GetObject(id)
	if err != nil {
		return nil
	}

	return object.(*player.Player)
}

func (pm *PlayerManager) GetPlayerByName(name string) *player.Player {
	var p *player.Player = nil
	pm.ObjectsRange(func(key, value interface{}) bool {
		if value.(*player.Player).Name == name {
			p = value.(*player.Player)
			return false
		}
		return true
	})

	return p
}

func (pm *PlayerManager) OnRemoveSession(sid zNet.SessionIdType) {

}

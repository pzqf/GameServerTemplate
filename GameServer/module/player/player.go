package player

import "github.com/pzqf/zEngine/zObject"

type Player struct {
	zObject.BaseObject
	Name string
}

func NewPlayer() *Player {
	return &Player{}
}

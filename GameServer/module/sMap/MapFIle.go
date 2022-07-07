package sMap

import (
	"encoding/json"

	"github.com/pzqf/zEngine/zNavigationMap"
)

type MapData struct {
	MapId   int                     `json:"map_id"`
	MapName string                  `json:"map_name"`
	MaxX    int                     `json:"max_x"`
	MaxY    int                     `json:"max_y"`
	Grids   [][]zNavigationMap.Grid `json:"grids"`
}

func LoadMap(mapFile string) (MapData, error) {
	var mapData MapData
	err := json.Unmarshal([]byte(mapFile), &mapData)
	if err != nil {
		return MapData{}, err
	}
	return mapData, nil
}

package sMap

import (
	"github.com/pzqf/zEngine/zNavigationMap"
	"github.com/pzqf/zEngine/zObject"
)

type SMap struct {
	zObject.BaseObject
	zNavigationMap.NavigationMap
	MapName string
}

func NewMap(mapFile string) *SMap {
	mapData, err := LoadMap(mapFile)
	if err != nil {
		return nil
	}
	m := SMap{
		MapName: mapData.MapName,
	}
	m.SetId(1)
	m.NavigationMap = zNavigationMap.NewNavigationMap(mapData.MaxX, mapData.MaxY, 1)
	for _, x := range mapData.Grids {
		for _, v := range x {
			err = m.NavigationMap.AddGrid(v)
			if err != nil {
				return nil
			}
		}
	}

	return &m
}

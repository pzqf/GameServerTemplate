package GameConfig

import (
	"github.com/pzqf/zUtil/zDataConv"
	"github.com/xuri/excelize/v2"
)

type MapConfig struct {
	MapId   int
	MapName string
	MapFile string
}

var MapConfigData []MapConfig

func LoadMapConfig(excelFile string) error {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	for k, row := range rows {
		if k == 0 {
			continue
		}

		m := MapConfig{}
		m.MapId, _ = zDataConv.String2Int(row[0])
		m.MapName = row[1]
		m.MapFile = row[2]
		MapConfigData = append(MapConfigData, m)
	}

	//fmt.Println(MapConfigData)
	return nil
}

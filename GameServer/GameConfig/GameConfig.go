package GameConfig

import "path"

// MapConfigData 地图配置数据
var MapConfigData []MapConfig

func LoadGameConfig(dirStr string) {
	err := LoadMapConfig(path.Join(dirStr, "test_game_config.xlsx"))
	if err != nil {
		return
	}
}

package GameConfig

import "path"

func LoadGameConfig(dirStr string) {
	err := LoadMapConfig(path.Join(dirStr, "test_game_config.xlsx"))
	if err != nil {
		return
	}
}

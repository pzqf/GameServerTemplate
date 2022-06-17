module ServerTemplate

go 1.18

require (
	github.com/pelletier/go-toml v1.9.5
	github.com/pzqf/zEngine v0.0.1
)

require (
	github.com/panjf2000/ants v1.3.0 // indirect
	github.com/pzqf/zUtil v0.0.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/pzqf/zEngine => ../zEngine

replace github.com/pzqf/zUtil => ../zUtil

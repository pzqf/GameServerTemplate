package monster

import (
	"github.com/pzqf/zEngine/zObject"
	"github.com/pzqf/zEngine/zScript"
)

type Monster struct {
	zScript.ScriptHolder
	zObject.BaseObject
	Name string
}

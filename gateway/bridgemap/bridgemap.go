package bridgemap

import (
	"github.com/t0mer/matterbridge/bridge"
)

var (
	FullMap           = map[string]bridge.Factory{}
	UserTypingSupport = map[string]struct{}{}
)

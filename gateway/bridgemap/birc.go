// +build !noirc

package bridgemap

import (
	birc "github.com/t0mer/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}

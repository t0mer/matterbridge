// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/t0mer/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}

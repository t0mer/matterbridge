// +build !nomumble

package bridgemap

import (
	bmumble "github.com/t0mer/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}

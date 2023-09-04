// +build !nosteam

package bridgemap

import (
	bsteam "github.com/t0mer/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}

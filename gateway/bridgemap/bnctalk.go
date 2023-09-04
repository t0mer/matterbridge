// +build !nonctalk

package bridgemap

import (
	btalk "github.com/t0mer/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}

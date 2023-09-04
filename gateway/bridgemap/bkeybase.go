// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/t0mer/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}

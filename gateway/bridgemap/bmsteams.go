// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/t0mer/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}

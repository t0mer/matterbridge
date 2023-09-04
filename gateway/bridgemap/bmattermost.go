// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/t0mer/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}

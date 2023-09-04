// +build !nozulip

package bridgemap

import (
	bzulip "github.com/t0mer/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}

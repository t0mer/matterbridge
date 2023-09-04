// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/t0mer/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}

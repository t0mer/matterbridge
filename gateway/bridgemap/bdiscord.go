// +build !nodiscord

package bridgemap

import (
	bdiscord "github.com/t0mer/matterbridge/bridge/discord"
)

func init() {
	FullMap["discord"] = bdiscord.New
	UserTypingSupport["discord"] = struct{}{}
}

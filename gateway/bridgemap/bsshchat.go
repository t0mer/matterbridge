// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/t0mer/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}

// +build !notelegram

package bridgemap

import (
	btelegram "github.com/t0mer/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}

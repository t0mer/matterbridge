// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/t0mer/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}

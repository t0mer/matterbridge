// +build !nowhatsapp
// +build !whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/t0mer/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}

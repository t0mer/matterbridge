// +build !noapi

package bridgemap

import (
	"github.com/t0mer/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}

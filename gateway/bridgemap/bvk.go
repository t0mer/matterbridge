// +build !novk

package bridgemap

import (
	bvk "github.com/t0mer/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}

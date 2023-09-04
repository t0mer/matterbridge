// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/t0mer/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}

//go:build !noharmony
// +build !noharmony

package bridgemap

import (
	bharmony "github.com/t0mer/matterbridge/bridge/harmony"
)

func init() {
	FullMap["harmony"] = bharmony.New
}

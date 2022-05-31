package vcs

import (
	"runtime/debug"
	"strconv"
)

type Vcs struct {
	Vcs      string
	Revision string
	Time     string
	Modified bool
}

// ReadInfo returns the vcs information embedded in the running binary.
// The information is available only in binaries built with module support.
// And for go 1.18+.
func ReadInfo() (*Vcs, bool) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, false
	}

	return FromBuildInfo(bi)
}

// FromBuildInfo returns the vcs information embedded in the provided BuildInfo object.
// The information is available only in binaries built with module support.
// And for go 1.18+.
func FromBuildInfo(info *debug.BuildInfo) (*Vcs, bool) {

	var vcs Vcs
	for _, s := range info.Settings {
		// Look out for vcs keys as defined at:
		// https://github.com/golang/go/blob/db19b42ca8771c25aa09e3747812f0229d44e75c/src/cmd/go/internal/load/pkg.go#L2450-L2458
		switch s.Key {
		case "vcs.revision":
			vcs.Revision = s.Value
		case "vcs.time":
			vcs.Time = s.Value
		case "vcs.modified":
			b, err := strconv.ParseBool(s.Value)
			if err != nil {
				return nil, false
			}
			vcs.Modified = b
		case "vcs":
			vcs.Vcs = s.Value
		}
	}

	// If the main "vcs" key is not set in BuildInfo,
	// discard all the other vcs keys and return nil.
	if vcs.Vcs == "" {
		return nil, false
	}

	return &vcs, true
}

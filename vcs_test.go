package vcs

import (
	"reflect"
	"runtime/debug"
	"testing"
)

// TestVcs tests Vcs extraction funcionality.
func TestVcs(t *testing.T) {
	tests := []struct {
		input *debug.BuildInfo
		want  *Vcs
		ok    bool
	}{
		{
			input: &debug.BuildInfo{
				Settings: []debug.BuildSetting{
					{Key: "a", Value: "b"},
					{Key: "p", Value: "g"},
					{Key: "d", Value: "t"},
					{Key: "k", Value: "e"},
				},
			},
			want: nil,
			ok:   false,
		},
		{
			input: &debug.BuildInfo{
				Settings: []debug.BuildSetting{
					{Key: "a", Value: "b"},
					{Key: "vcs", Value: "git"},
					{Key: "d", Value: "t"},
					{Key: "k", Value: "e"},
				},
			},
			want: &Vcs{Vcs: "git"},
			ok:   true,
		},
		{
			input: &debug.BuildInfo{
				Settings: []debug.BuildSetting{
					{Key: "a", Value: "b"},
					{Key: "vcs.modified", Value: "true"},
					{Key: "vcs.time", Value: "tttt"},
					{Key: "vcs.revision", Value: "rrrr"},
				},
			},
			want: nil,
			ok:   false,
		},
		{
			input: &debug.BuildInfo{
				Settings: []debug.BuildSetting{
					{Key: "a", Value: "b"},
					{Key: "vcs", Value: "git"},
					{Key: "vcs.modified", Value: "true"},
					{Key: "vcs.time", Value: "tttt"},
					{Key: "vcs.revision", Value: "rrrr"},
				},
			},
			want: &Vcs{Vcs: "git", Modified: true, Time: "tttt", Revision: "rrrr"},
			ok:   true,
		},
		{
			input: &debug.BuildInfo{
				Settings: []debug.BuildSetting{
					{Key: "a", Value: "b"},
					{Key: "vcs", Value: "git"},
					{Key: "vcs.modified", Value: "false"},
					{Key: "vcs.time", Value: "tttt"},
					{Key: "vcs.revision", Value: "rrrr"},
				},
			},
			want: &Vcs{Vcs: "git", Modified: false, Time: "tttt", Revision: "rrrr"},
			ok:   true,
		},
		{
			input: &debug.BuildInfo{
				Settings: []debug.BuildSetting{
					{Key: "a", Value: "b"},
					{Key: "vcs", Value: "git"},
					{Key: "vcs.modified", Value: "invalid"},
					{Key: "vcs.time", Value: "tttt"},
					{Key: "vcs.revision", Value: "rrrr"},
				},
			},
			want: nil,
			ok:   false,
		},
	}

	for _, tc := range tests {
		v, ok := FromBuildInfo(tc.input)

		if !reflect.DeepEqual(tc.want, v) {
			t.Fatalf("expected: %v, got: %v", tc.want, v)
		}
		if tc.ok != ok {
			t.Fatalf("expected: %v, got: %v", tc.ok, ok)
		}
	}
}

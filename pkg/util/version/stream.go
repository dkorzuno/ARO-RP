package version

import "fmt"

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

type Stream struct {
	Version    *Version
	PullSpec   string
	MustGather string
}

// InstallStream describes stream we are defaulting to for all new clusters
var InstallStream = Stream{
	Version:    NewVersion(4, 4, 10),
	PullSpec:   "quay.io/openshift-release-dev/ocp-release@sha256:0d1ffca302ae55d32574b38438c148d33c2a8a05c8daf97eeb13e9ab948174f7",
	MustGather: "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:06ae5b7e36f23eb2e5ae5826499978de4e0124e33938c2e532ed73563b1f7c14",
}

// Streams describes list of streams we support for upgrades
var (
	Streams = []Stream{
		InstallStream,
		{
			Version:    NewVersion(4, 3, 27),
			PullSpec:   "quay.io/openshift-release-dev/ocp-release@sha256:a2bdd3b4516e05760d01e2589fc0866f7386c1c10c866b29fea137067e76f2ae",
			MustGather: "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:78a3e629ec24ec56b6da90d24eb63ee06a47be2ec82833c66cb5a02aa6a7cc92",
		},
	}
)

// GetStream return matching stream, used to upgrade cluster.
func GetStream(v *Version) (*Stream, error) {
	for _, s := range Streams {
		if s.Version.V[0] == v.V[0] &&
			s.Version.V[1] == v.V[1] {

			// we DO NOT upgrade if CVO version is already higher
			if !v.Lt(s.Version) {
				return nil, fmt.Errorf("not upgrading: cvo desired version is %s", v)
			}

			return &s, nil
		}
	}
	return nil, fmt.Errorf("stream for %s not found", v)
}
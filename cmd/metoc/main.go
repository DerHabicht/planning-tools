package main

import (
	"fmt"

	"github.com/derhabicht/planning-tools/internal/cmd/metoc"
)

// BaseVersion is the SemVer-formatted string that defines the current version of metoc.
// Build information will be added at compile-time.
const BaseVersion = "1.0.0"

// BuildTime is a timestamp of when the build is run. This variable is set at compile-time.
var BuildTime string

// GitRevision is the current Git commit ID. If the tree is dirty at compile-time, an "x-" is prepended to the hash.
// This variable is set at compile-time.
var GitRevision string

// GitBranch is the name of the active Git branch at compile-time. This variable is set at compile-time.
var GitBranch string

func main() {
	version := fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)
	metoc.Execute(version)
}

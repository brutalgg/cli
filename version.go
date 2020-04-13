package cli

import "fmt"

// Version is current version of this library.
var Version = v{1, 0, 0}

// v holds the version of this library.
type v struct {
	Major, Minor, Patch int
}

func (v v) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
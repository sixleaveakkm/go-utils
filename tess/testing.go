/*
	tess is a package contains helper functions used while testing
*/

package tess

import (
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// SkipCI skips test case if environment has `CI` set to true ("true", "TRUE", "True", "1").
// Example:
//
//	function TestFoo(t *testing.T) {
//	    SkipCI(t)
//	    // do some test you want to be skipped in CI (only local)
//	}
func SkipCI(t *testing.T) {
	t.Helper()
	if isCI() {
		t.Skip("Skipping testing in CI environment")
	}
}

func isCI() bool {
	return isTrue(os.Getenv("CI"))
}

func isTrue(s string) bool {
	switch s {
	case "true", "True", "TRUE", "1":
		return true
	default:
		return false
	}
}

// IsSameTime test two given time are the same, with an optional parameter to set the time drift threshold (default set to 100ms)
func IsSameTime(tm1 time.Time, tm2 time.Time, drift ...time.Duration) bool {
	d := 100 * time.Millisecond
	if len(drift) != 0 {
		d = drift[0]
	}
	diff := tm1.Sub(tm2)
	abs := diff.Abs()
	return abs < d
}

// goCmpTimeComparer helper var for go-cmp.Diff function.
var goCmpTimeComparer = cmp.Comparer(func(x, y time.Time) bool {
	return IsSameTime(x, y)
})

// Diff is a warp of go-cmp diff with default optional comparer.
// Including a custom time comparer allows 100ms diff.
func Diff(x, y any) string {
	return cmp.Diff(x, y, goCmpTimeComparer)
}

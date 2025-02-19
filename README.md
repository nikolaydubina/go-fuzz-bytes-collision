> Issue: https://github.com/golang/go/issues/71831

Example showing that `[]byte` arguments in Go fuzz testing are colliding and unreliable without copying them to new slices first.

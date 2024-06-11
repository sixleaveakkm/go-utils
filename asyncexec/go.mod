module github.com/sixleaveakkm/go-utils/asyncexec

go 1.19

require (
	github.com/sixleaveakkm/go-utils v1.2.0-alpha
	github.com/sixleaveakkm/go-utils/errz v0.1.0
	github.com/sixleaveakkm/go-utils/ptr v1.0.1
	github.com/sixleaveakkm/go-utils/toy v0.0.0-00010101000000-000000000000
	golang.org/x/time v0.5.0
)

replace (
	github.com/sixleaveakkm/go-utils/toy => ../toy
)

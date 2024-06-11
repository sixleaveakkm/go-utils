module github.com/sixleaveakkm/go-utils

go 1.19

require (
	github.com/stretchr/testify v1.9.0
	github.com/sixleaveakkm/go-utils/errz v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/sixleaveakkm/go-utils/errz => ./errz
)
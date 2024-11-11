module github.com/apieasy/gson

go 1.23

require (
	github.com/stretchr/testify v1.9.0
	github.com/tidwall/gjson v1.18.0
	github.com/tidwall/sjson v1.2.5
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// https://pkg.go.dev/about#adding-a-package
// https://go.dev/ref/mod#go-mod-file-retract
retract [v0.1.0, v0.1.1]

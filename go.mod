module PVZ

go 1.23.2

replace PVZ/internal/generated-client => ./internal/generated-client

require PVZ/internal/generated-client v0.0.0-00010101000000-000000000000

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

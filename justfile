set positional-arguments
set export

@default:
	just --list --unsorted

#
# Project settings
#
BinaryName := "./example-ent-crdb"
ModuleName := "github.com/kofalt/example-ent-crdb"

#
# Developer
#

# Run server
run: gen
	go run -v .

# Pre-commit sanity checks
pre: gen fmt lint

# Serve local documentation
doc: get-pkgsite
	@echo "Browse to http://localhost:6060/{{ModuleName}}"
	pkgsite -http localhost:6060

# ORM CLI
@ent *args='':
	# This matches the ent CLI against the version in go.mod
	# Otherwise go/bin versions could drift
	go run -mod=mod entgo.io/ent/cmd/ent "$@"


#
# Util
#


# Update generated code
gen:
	go generate ./ent

# Format code
fmt:
	gofmt -w .

# Lint code
lint: get-linter
	golangci-lint run

#
# Tools
#

@get-linter:
	hash golangci-lint || (echo "Installing golint..."; go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1)

@get-pkgsite:
	hash pkgsite || (echo "Installing pkgsite..."; go install -v golang.org/x/pkgsite/cmd/pkgsite@latest)

#
# CI
#

# CI target
ci: check-fmt lint release

# Format code, exiting on error
@check-fmt:
	test -z "$(gofmt -l .)" || (echo "One or more lines need formatting"; exit 1)

# A CI runner can set these
# Could modify to use env_var_or_default
BuildHash  := "dev"
BuildDate  := "dev"

# Golang variable inject syntax
VersionFlags := "-X main.BuildHash=" + BuildHash + " -X main.BuildDate=" + BuildDate

# Cut binary
release: gen
	rm -f "$BinaryName"
	go build -ldflags "$VersionFlags" -v -o "$BinaryName" .
	strip --strip-debug --strip-unneeded "$BinaryName"


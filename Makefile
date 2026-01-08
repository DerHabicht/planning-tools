SRC = $(shell find . -name "*.go")

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

.PHONY: all
all: bin/plancal bin/metoc

bin/plancal: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/plancal cmd/plancal/main.go

bin/metoc: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/metoc cmd/metoc/main.go

bin/s5: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/s5 cmd/s5/main.go

.PHONY: install
install: bin/plancal bin/metoc
	go run build/install.go $(CURDIR)
	cp bin/plancal ${HOME}/.local/bin/
	cp bin/metoc ${HOME}/.local/bin/
	cp bin/s5 ${HOME}/.local/bin/

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -rf bin/

CGO_ENABLED ?= 0
GIT_COMMIT ?= $$(git rev-parse HEAD)
GIT_REPOSITORY ?= $$(git config --get remote.origin.url)

build:
	go build \
        -ldflags \
            "-X github.com/SAP/jenkins-library/cmd.GitCommit=${GIT_COMMIT} \
            -X github.com/SAP/jenkins-library/pkg/log.LibraryRepository=${GIT_REPOSITORY} \
            -X github.com/SAP/jenkins-library/pkg/telemetry.LibraryRepository=${GIT_REPOSITORY}" \
        -o piper

install:
	cp piper ~/go/bin/

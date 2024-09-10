IMAGE_REPOSITORY = ecr.elevenz.com/software
VERSION = v0.0.1

build: MAKE_IMAGE ?= ${IMAGE_REPOSITORY}/shell-runner:${VERSION}
build:
	@go mod init mysql-runner
	@go mod tidy
	@docker buildx build -t ${MAKE_IMAGE} --platform linux/amd64,linux/arm64 -f Dockerfile --push .
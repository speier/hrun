VERSION=v0.1.4

release:
	@git tag -a ${VERSION} -m "Release ${VERSION}" && git push origin ${VERSION}
	@goreleaser --rm-dist

VERSION=v0.1.1

release:
	@git tag -a ${VERSION} -m "Release ${VERSION}" && git push origin ${VERSION}
	@goreleaser --rm-dist

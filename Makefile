VERSION=v0.2.1
MODULE=github.com/srmullen/godraw-lib

# https://go.dev/doc/modules/publishing
publish:
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)
	GOPROXY=proxy.golang.org go list -m $(MODULE)@$(VERSION) 